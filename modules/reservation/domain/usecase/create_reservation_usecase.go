package usecase

import (
	"bytes"
	"context"
	"github.com/skip2/go-qrcode"
	"image"
	"trekkstay/core"
	"trekkstay/modules/reservation/domain/entity"
	"trekkstay/pkgs/s3"
)

type CreateReservationUseCase interface {
	ExecuteCreateReservation(ctx context.Context, reservation *entity.ReservationEntity) (*entity.ReservationEntity, error)
}

type createReservationUseCaseImpl struct {
	roomReaderRepo        HotelRoomReaderRepository
	reservationWriterRepo ReservationWriterRepository
	uploadHandler         s3.UploadHandler
}

var _ CreateReservationUseCase = (*createReservationUseCaseImpl)(nil)

func NewCreateReservationUseCase(
	roomReaderRepo HotelRoomReaderRepository,
	reservationWriterRepo ReservationWriterRepository,
	uploadHandler s3.UploadHandler,
) CreateReservationUseCase {
	return &createReservationUseCaseImpl{
		roomReaderRepo:        roomReaderRepo,
		reservationWriterRepo: reservationWriterRepo,
		uploadHandler:         uploadHandler,
	}
}

func generateQRCode(text string) (image.Image, error) {
	// Generate QR code
	qr, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	// Create image from QR code bytes
	img, _, err := image.Decode(bytes.NewReader(qr))
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (useCase createReservationUseCaseImpl) ExecuteCreateReservation(ctx context.Context,
	reservation *entity.ReservationEntity) (*entity.ReservationEntity, error) {
	// Find room
	room, err := useCase.roomReaderRepo.FindHotelRoomByCondition(ctx,
		map[string]interface{}{"id": reservation.RoomID})

	if err != nil {
		return nil, err
	}

	// Retrieve room information
	(*reservation).Room.HotelID = (*room).HotelID
	(*reservation).Room.Type = (*room).Type
	(*reservation).Room.OriginalPrice = (*room).OriginalPrice
	(*reservation).Room.Images = entity.MediaJSON((*room).Images)
	(*reservation).Room.BookingPrice = (*room).OriginalPrice - ((*room).OriginalPrice * (*room).DiscountRate / 100)

	(*reservation).Status = "UPCOMING"
	(*reservation).UserID = ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserID()
	(*reservation).TotalPrice = (*reservation).Room.BookingPrice * (*reservation).Quantity

	if err := useCase.reservationWriterRepo.InsertReservation(ctx, reservation); err != nil {
		return nil, err
	}

	qrCode, _ := generateQRCode(reservation.ID)
	url, err := useCase.uploadHandler.UploadImageToS3(qrCode)
	(*reservation).QRCodeURL = *url

	if err := useCase.reservationWriterRepo.UpdateReservation(ctx, *reservation); err != nil {
		return nil, err
	}

	return reservation, nil
}
