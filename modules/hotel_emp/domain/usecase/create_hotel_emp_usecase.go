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
		log.JsonLogger.Error("ExecuteCreateHotelEmp.find_hotel_emp_by_id",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrEmpNotFound(err)
	}

	// Check if requester is owner
	if requester.Role != constant.OwnerRole {
		log.JsonLogger.Error("ExecuteCreateHotelEmp.find_hotel_owner_by_id",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrPermissionDenied(errors.New("you are not owner of this hotel"))
	}

	// Check if hotel emp already exists
	_, err = useCase.readerRepo.FindHotelEmpByCondition(ctx, map[string]interface{}{
		"email": hotelEmpEntity.Email,
	})
	if err == nil {
		log.JsonLogger.Error("ExecuteCreateHotelEmp.hotel_emp_already_exists",
			slog.String("error", errors.New("email already exists").Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorEmailAlreadyExists(nil)
	}

	// Check if hotel emp already exists
	_, err = useCase.readerRepo.FindHotelEmpByCondition(ctx, map[string]interface{}{
		"phone": hotelEmpEntity.Phone,
	})
	if err == nil {
		log.JsonLogger.Error("ExecuteCreateHotelEmp.hotel_emp_already_exists",
			slog.String("error", errors.New("phone already exists").Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorPhoneAlreadyExists(nil)
	}

	// Random password and hash password
	randomPassword, _ := utils.GenerateRandomPassword(8)
	hashedPassword, err := useCase.hashAlgo.HashAndSalt([]byte(randomPassword))
	if err != nil {
		log.JsonLogger.Error("ExecuteCreateHotelEmp.hash_password",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorHashPassword(err)
	}

	// Assign password, hotel id and role
	hotelEmpEntity.Password = hashedPassword
	hotelEmpEntity.HotelID = requester.HotelID
	hotelEmpEntity.Role = constant.EmpRole

	if hotelEmpEntity.HotelID == "" {
		log.JsonLogger.Error("ExecuteCreateHotelEmp.hotel_id_is_empty",
			slog.String("error", "hotel id is empty"),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrNoHotelToAssignEmp(nil)
	}

	// Send email
	go func() {
		err = useCase.mailer.SendMail(hotelEmpEntity.Email, "Create hotel employee", utils.GetWorkingDirectory()+
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

	// Create hotel employee
	if err := useCase.writerRepo.InsertHotelEmp(ctx, hotelEmpEntity); err != nil {
		log.JsonLogger.Error("ExecuteCreateHotelEmp.create_hotel_emp",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorInternalServerError(err)
	}

	return nil
}
