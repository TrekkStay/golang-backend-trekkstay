package usecase

import (
	"context"
	"errors"
	"log/slog"
	"trekkstay/core"
	"trekkstay/modules/hotel_emp/constant"
	"trekkstay/modules/hotel_emp/domain/entity"
	"trekkstay/pkgs/log"
	"trekkstay/utils"
)

type CreateHotelEmpUseCase interface {
	ExecuteCreateHotelEmp(ctx context.Context, hotelEmpEntity entity.HotelEmpEntity) error
}

type hotelEmpUseCaseImpl struct {
	mailer     Mailer
	hashAlgo   HashAlgo
	readerRepo hotelEmpReaderRepository
	writerRepo hotelEmpWriterRepository
}

var _ CreateHotelEmpUseCase = (*hotelEmpUseCaseImpl)(nil)

func NewCreateHotelEmpUseCase(mailer Mailer, hashAlgo HashAlgo, readerRepo hotelEmpReaderRepository,
	writerRepo hotelEmpWriterRepository) CreateHotelEmpUseCase {
	return &hotelEmpUseCaseImpl{
		mailer:     mailer,
		hashAlgo:   hashAlgo,
		readerRepo: readerRepo,
		writerRepo: writerRepo,
	}
}

func (useCase hotelEmpUseCaseImpl) ExecuteCreateHotelEmp(ctx context.Context, hotelEmpEntity entity.HotelEmpEntity) error {
	// Find hotel emp by id from token
	requester, err := useCase.readerRepo.FindHotelEmpByCondition(ctx, map[string]interface{}{
		"id": ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserID(),
	})
	if err != nil {
		return handleError("find_hotel_emp_by_id", "ExecuteCreateHotelEmp", err, ctx)
	}

	if err := checkRequesterRole(*requester); err != nil {
		return err
	}

	if err := checkExistingEmail(useCase, ctx, hotelEmpEntity.Email, "email_already_exists"); err != nil {
		return err
	}

	if err := checkExistingPhone(useCase, ctx, hotelEmpEntity.Phone, "phone_already_exists"); err != nil {
		return err
	}

	randomPassword, hashedPassword, err := generateAndHashPassword(useCase)
	if err != nil {
		return handleError("hash_password", "ExecuteCreateHotelEmp", err, ctx)
	}

	hotelEmpEntity.Password = hashedPassword
	hotelEmpEntity.HotelID = requester.HotelID
	hotelEmpEntity.Role = constant.EmpRole

	if hotelEmpEntity.HotelID == "" {
		return handleError("hotel_id_is_empty", "ExecuteCreateHotelEmp", errors.New("hotel id is empty"), ctx)
	}

	if err := useCase.writerRepo.InsertHotelEmp(ctx, hotelEmpEntity); err != nil {
		return handleError("create_hotel_emp", "ExecuteCreateHotelEmp", err, ctx)
	}

	sendEmail(useCase, ctx, hotelEmpEntity.Email, randomPassword)

	return nil
}

func handleError(errorType, funcName string, err error, ctx context.Context) error {
	log.JsonLogger.Error(funcName+"."+errorType,
		slog.String("error", err.Error()),
		slog.String("request_id", ctx.Value("X-Request-ID").(string)),
	)

	return err
}

func checkRequesterRole(requester entity.HotelEmpEntity) error {
	if requester.Role != constant.OwnerRole {
		return constant.ErrPermissionDenied(errors.New("you are not owner of this hotel"))
	}
	return nil
}

func checkExistingEmail(useCase hotelEmpUseCaseImpl, ctx context.Context, email string, errorType string) error {
	oldEmp, err := useCase.readerRepo.FindHotelEmpByCondition(ctx, map[string]interface{}{
		"email": email,
	})
	if oldEmp != nil {
		return constant.ErrorEmailAlreadyExists(err)
	}
	return nil
}

func checkExistingPhone(useCase hotelEmpUseCaseImpl, ctx context.Context, phone string, errorType string) error {
	oldEmp, err := useCase.readerRepo.FindHotelEmpByCondition(ctx, map[string]interface{}{
		"phone": phone,
	})
	if oldEmp != nil {
		return constant.ErrorPhoneAlreadyExists(err)
	}
	return nil
}

func generateAndHashPassword(useCase hotelEmpUseCaseImpl) (string, string, error) {
	randomPassword, _ := utils.GenerateRandomPassword(8)
	hashedPassword, err := useCase.hashAlgo.HashAndSalt([]byte(randomPassword))
	return randomPassword, hashedPassword, err
}

func sendEmail(useCase hotelEmpUseCaseImpl, ctx context.Context, email, randomPassword string) {
	go func() {
		err := useCase.mailer.SendMail(email, "Create hotel employee", utils.GetWorkingDirectory()+
			"/templates/send_password_to_emp.html", map[string]interface{}{
			"password": randomPassword,
		})

		if err != nil {
			log.JsonLogger.Error("ExecuteCreateHotelEmp.send_mail",
				slog.Any("error", err.Error()),
				slog.String("request_id", ctx.Value("X-Request-ID").(string)),
			)
		}
	}()
}
