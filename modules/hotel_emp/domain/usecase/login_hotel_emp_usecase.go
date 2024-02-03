package usecase

import (
	"context"
	"log/slog"
	"trekkstay/api/middlewares/constant"
	"trekkstay/modules/hotel_emp/domain/entity"
	"trekkstay/pkgs/log"
)

type LoginHotelEmpUseCase interface {
	ExecuteLoginHotelEmp(ctx context.Context, hotelEmpEntity entity.HotelEmpEntity) (*entity.HotelEmpEntity, error)
}

type loginHotelEmpUseCaseImpl struct {
	tokenProvider    TokenProvider
	accessTokenTime  int
	refreshTokenTime int
	hashAlgo         HashAlgo
	readerRepo       hotelEmpReaderRepository
}

var _ LoginHotelEmpUseCase = (*loginHotelEmpUseCaseImpl)(nil)

func NewLoginHotelEmpUseCase(tokenProvider TokenProvider, accessTokenTime int, refreshTokenTime int,
	hashAlgo HashAlgo, readerRepo hotelEmpReaderRepository) LoginHotelEmpUseCase {
	return &loginHotelEmpUseCaseImpl{
		tokenProvider:    tokenProvider,
		accessTokenTime:  accessTokenTime,
		refreshTokenTime: refreshTokenTime,
		hashAlgo:         hashAlgo,
		readerRepo:       readerRepo,
	}
}

func (useCase loginHotelEmpUseCaseImpl) ExecuteLoginHotelEmp(ctx context.Context,
	hotelEmpEntity entity.HotelEmpEntity) (*entity.HotelEmpEntity, error) {
	// Check if emp exists
	emp, err := useCase.readerRepo.FindHotelEmpByCondition(ctx, map[string]interface{}{
		"email": hotelEmpEntity.Email,
	})
	if err != nil {
		log.JsonLogger.Error("ExecuteLoginHotelEmp.email_not_found",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return nil, constant.ErrEmailNotFound(err)
	}

	// Check if password is correct
	if err := useCase.hashAlgo.ComparePasswords(emp.Password, []byte(hotelEmpEntity.Password)); err != nil {
		log.JsonLogger.Error("ExecuteLoginHotelEmp.password_not_match",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return nil, constant.ErrEmailOrPasswordInvalid(err)
	}

	// Generate access token
	accessToken, err := useCase.tokenProvider.Generate(
		map[string]interface{}{
			"user_id": emp.ID,
			"role":    emp.Role,
		},
		useCase.accessTokenTime,
	)
	if err != nil {
		log.JsonLogger.Error("ExecLoginUser.generate_access_token",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return nil, constant.ErrInternal(err)
	}

	// Generate refresh token
	refreshToken, err := useCase.tokenProvider.Generate(
		map[string]interface{}{
			"user_id": emp.ID,
			"role":    emp.Role,
		},
		useCase.refreshTokenTime,
	)
	if err != nil {
		log.JsonLogger.Error("ExecLoginUser.generate_refresh_token",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return nil, constant.ErrInternal(err)
	}

	emp.AccessToken = accessToken["token"].(string)
	emp.RefreshToken = refreshToken["token"].(string)

	return emp, nil
}
