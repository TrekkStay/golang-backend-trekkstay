package usecase

import (
	"context"
	"errors"
	"log/slog"
	"trekkstay/modules/hotel_emp/constant"
	"trekkstay/modules/hotel_emp/domain/entity"
	"trekkstay/pkgs/log"
)

type CreateHotelOwnerUseCase interface {
	ExecuteCreateHotelOwner(ctx context.Context, hotelOwnerEntity entity.HotelEmpEntity) error
}

type hotelOwnerUseCaseImpl struct {
	hashAlgo   HashAlgo
	readerRepo hotelEmpReaderRepository
	writerRepo hotelEmpWriterRepository
}

var _ CreateHotelOwnerUseCase = (*hotelOwnerUseCaseImpl)(nil)

func NewCreateHotelOwnerUseCase(hashAlgo HashAlgo, readerRepo hotelEmpReaderRepository,
	writerRepo hotelEmpWriterRepository) CreateHotelOwnerUseCase {
	return &hotelOwnerUseCaseImpl{
		hashAlgo:   hashAlgo,
		readerRepo: readerRepo,
		writerRepo: writerRepo,
	}
}

func (useCase hotelOwnerUseCaseImpl) ExecuteCreateHotelOwner(ctx context.Context, hotelOwnerEntity entity.HotelEmpEntity) error {
	// Find hotel owner by email
	hotelOwner, err := useCase.readerRepo.FindHotelEmpByCondition(ctx, map[string]interface{}{
		"email": hotelOwnerEntity.Email,
	})
	if hotelOwner != nil {
		log.JsonLogger.Error("ExecuteCreateHotelOwner.email_already_exists",
			slog.Any("error", errors.New("email already exists")),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorEmailAlreadyExists(err)
	}

	// Find hotel owner by phone
	hotelOwner, err = useCase.readerRepo.FindHotelEmpByCondition(ctx, map[string]interface{}{
		"phone": hotelOwnerEntity.Phone,
	})
	if hotelOwner != nil {
		log.JsonLogger.Error("ExecuteCreateHotelOwner.phone_already_exists",
			slog.Any("error", errors.New("phone already exists")),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorPhoneAlreadyExists(err)
	}

	// Hash password
	hashedPassword, err := useCase.hashAlgo.HashAndSalt([]byte(hotelOwnerEntity.Password))
	if err != nil {
		log.JsonLogger.Error("ExecuteCreateHotelOwner.hash_password_error",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorHashPassword(err)
	}

	// Set hashed password. Role and contract
	hotelOwnerEntity.Password = hashedPassword
	hotelOwnerEntity.Role = constant.OwnerRole
	hotelOwnerEntity.Contract = constant.FullTimeContract

	// Create hotel owner
	if err := useCase.writerRepo.InsertHotelEmp(ctx, hotelOwnerEntity); err != nil {
		log.JsonLogger.Error("ExecuteCreateHotelOwner.insert_hotel_owner_error",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorInternalServerError(err)
	}

	return nil
}
