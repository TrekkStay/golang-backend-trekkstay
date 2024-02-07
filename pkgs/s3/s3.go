package s3

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"image"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"
	"trekkstay/config/models"
	res "trekkstay/core/response"
)

const (
	maxPartSize = 5 * 1024 * 1024
	maxRetries  = 3
)

type UploadHandler struct {
	s3Config *models.S3Config
}

type Form struct {
	Files []*multipart.FileHeader `form:"files" binding:"required"`
}

func NewS3Upload(s3Config *models.S3Config) UploadHandler {
	return UploadHandler{s3Config: s3Config}
}

// HandleUploadMedia godoc
// @Summary      Upload media files
// @Description  Upload media files to the server and store them in AWS S3
// @Tags         Upload
// @Accept       multipart/form-data
// @Produce      json
// @Param        files formData file true "Files to upload"
// @Success      200 {object} res.SuccessResponse
// @Failure      400 {object} res.ErrorResponse
// @Failure      500 {object} res.ErrorResponse
// @Router       /upload/media [post]
func (h *UploadHandler) HandleUploadMedia(c *gin.Context) {
	// Parse form data to extract files
	form, err := parseFormData(c)
	if err != nil {
		panic(err)
	}

	// Initialize AWS session and S3 client
	svc := initializeS3Client(h.s3Config)

	// Upload files to S3 and retrieve URLs
	urls := uploadFilesToS3(form.Files, svc, h.s3Config)

	response := map[string]interface{}{
		"urls": urls,
	}
	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", response))
}

func (h *UploadHandler) UploadImageToS3(image image.Image) (*string, error) {
	// Create a new AWS session with the specified credentials
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(h.s3Config.S3Region),
		Credentials: credentials.NewStaticCredentials(h.s3Config.S3AccessKey, h.s3Config.S3SecretKey, ""),
	})
	if err != nil {
		return nil, err
	}

	// Create an S3 client
	s3Client := s3.New(sess)

	// Create a buffer to store the PNG data
	var pngData bytes.Buffer
	err = png.Encode(&pngData, image)
	if err != nil {
		return nil, err
	}

	imageName := "qr-code/" + uuid.New().String() + ".png"

	// Upload the PNG data to S3
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(h.s3Config.S3Bucket),
		Key:    aws.String(imageName), // S3 object key
		Body:   bytes.NewReader(pngData.Bytes()),
	})
	if err != nil {
		return nil, err
	}

	url := h.s3Config.S3CloudFront + imageName
	return &url, nil
}

// parseFormData parses the form data from the request.
func parseFormData(c *gin.Context) (*Form, error) {
	var form Form
	if err := c.ShouldBind(&form); err != nil {
		return nil, err
	}
	return &form, nil
}

// initializeS3Client initializes the AWS session and S3 client.
func initializeS3Client(s3Config *models.S3Config) *s3.S3 {
	s3Credentials := credentials.NewStaticCredentials(s3Config.S3AccessKey, s3Config.S3SecretKey, "")
	_, err := s3Credentials.Get()
	if err != nil {
		panic(err)
	}

	cfg := aws.NewConfig().WithRegion(s3Config.S3Region).WithCredentials(s3Credentials)
	sess, err := session.NewSession(cfg)
	if err != nil {
		panic(err)
	}

	return s3.New(sess, cfg)
}

// uploadFilesToS3 uploads files to S3 and returns their URLs.
func uploadFilesToS3(files []*multipart.FileHeader, svc *s3.S3, s3Config *models.S3Config) []string {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var urls []string

	for _, file := range files {
		wg.Add(1)
		go func(file *multipart.FileHeader) {
			defer wg.Done()

			fileBytes, err := readFile(file)
			if err != nil {
				panic(err)
			}

			fileName := generateFilename(file.Filename)
			key := "/media/" + fileName
			contentType := http.DetectContentType(fileBytes)

			input := &s3.CreateMultipartUploadInput{
				Bucket:      aws.String(s3Config.S3Bucket),
				Key:         aws.String(key),
				ContentType: aws.String(contentType),
			}

			resp, err := svc.CreateMultipartUpload(input)
			if err != nil {
				panic(err)
			}

			completedParts := uploadFilePartsToS3(svc, resp, fileBytes, file.Size)

			_, err = completeMultipartUpload(svc, resp, completedParts)
			if err != nil {
				panic(err)
			}

			mu.Lock()
			urls = append(urls, s3Config.S3CloudFront+key)
			mu.Unlock()
		}(file)
	}

	wg.Wait()
	return urls
}

// uploadFilePartsToS3 uploads parts of a file to S3 and returns the completed parts.
func uploadFilePartsToS3(svc *s3.S3, resp *s3.CreateMultipartUploadOutput, fileBytes []byte, fileSize int64) []*s3.CompletedPart {
	var completedParts []*s3.CompletedPart
	for partNumber, offset := 1, int64(0); offset < fileSize; partNumber++ {
		partLength := min(int64(maxPartSize), fileSize-offset)

		completedPart, err := uploadPart(svc, resp, fileBytes[offset:offset+partLength], partNumber)
		if err != nil {
			panic(err)
		}

		completedParts = append(completedParts, completedPart)
		offset += partLength
	}
	return completedParts
}

func readFile(file *multipart.FileHeader) ([]byte, error) {
	openedFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer func(openedFile multipart.File) {
		err := openedFile.Close()
		if err != nil {
			fmt.Println("failed to close file")
		}
	}(openedFile)

	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

func generateFilename(originalFilename string) string {
	ext := strings.ToLower(strings.Split(originalFilename, ".")[1])
	return uuid.New().String() + "." + ext
}

func uploadPart(svc *s3.S3, resp *s3.CreateMultipartUploadOutput, fileBytes []byte, partNumber int) (*s3.CompletedPart, error) {
	partInput := &s3.UploadPartInput{
		Body:          bytes.NewReader(fileBytes),
		Bucket:        resp.Bucket,
		Key:           resp.Key,
		PartNumber:    aws.Int64(int64(partNumber)),
		UploadId:      resp.UploadId,
		ContentLength: aws.Int64(int64(len(fileBytes))),
	}

	tryNum := 1
	for tryNum <= maxRetries {
		uploadResult, err := svc.UploadPart(partInput)
		if err == nil {
			return &s3.CompletedPart{
				ETag:       uploadResult.ETag,
				PartNumber: aws.Int64(int64(partNumber)),
			}, nil
		}

		var uploadErr awserr.Error
		if errors.As(err, &uploadErr) {
			return nil, uploadErr
		}

		tryNum++
	}

	return nil, errors.New("failed to upload part")
}

func completeMultipartUpload(svc *s3.S3, resp *s3.CreateMultipartUploadOutput,
	completedParts []*s3.CompletedPart) (*s3.CompleteMultipartUploadOutput, error) {
	completeInput := &s3.CompleteMultipartUploadInput{
		Bucket:   resp.Bucket,
		Key:      resp.Key,
		UploadId: resp.UploadId,
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: completedParts,
		},
	}
	return svc.CompleteMultipartUpload(completeInput)
}
