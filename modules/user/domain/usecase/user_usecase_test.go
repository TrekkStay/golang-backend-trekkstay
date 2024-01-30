package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"trekkstay/core"
	"trekkstay/modules/user/domain/entity"
)

func TestCreateUserUseCase(t *testing.T) {
	userReaderRepo := mockUserReaderRepository{}
	userWriterRepo := mockUserWriterRepository{}
	hashAlgo := mockHashAlgo{}

	useCase := NewCreateUserUseCase(hashAlgo, userReaderRepo, userWriterRepo)

	var ctx = context.WithValue(context.Background(), "X-Request-ID", "1234567890")
	t.Run("create user successfully", func(t *testing.T) {
		err := useCase.ExecCreateUser(ctx, entity.UserEntity{
			FullName: "Test User",
			Email:    "testuseraasd@example.com",
			Phone:    "1234567890",
			Status:   entity.ACTIVE.Value(),
			OTP:      "123456",
			Password: "password",
		})

		assert.Nil(t, err)
	})

	t.Run("email already exist", func(t *testing.T) {
		err := useCase.ExecCreateUser(ctx, entity.UserEntity{
			FullName: "Test User",
			Email:    "existedemail@example.com",
			Phone:    "1234567890",
			Status:   entity.ACTIVE.Value(),
			OTP:      "123456",
			Password: "password",
		})

		assert.NotNil(t, err)
	})
}

func TestLoginUserUseCase(t *testing.T) {
	userReaderRepo := mockUserReaderRepository{}
	tokenProvider := mockTokenProvider{}
	hashAlgo := mockHashAlgo{}

	useCase := NewLoginUserUseCase(tokenProvider, 1, 1, hashAlgo, userReaderRepo)

	var ctx = context.WithValue(context.Background(), "X-Request-ID", "1234567890")
	t.Run("login user successfully", func(t *testing.T) {
		userEntity, err := useCase.ExecLoginUser(ctx, entity.UserEntity{
			Email:    "existedemail@example.com",
			Password: "password",
		})

		assert.Nil(t, err)
		assert.NotNil(t, userEntity)
	})
}

func TestChangePasswordUseCase(t *testing.T) {
	userReaderRepo := mockUserReaderRepository{}
	userWriterRepo := mockUserWriterRepository{}
	hashAlgo := mockHashAlgo{}

	useCase := NewChangePasswordUseCase(hashAlgo, userReaderRepo, userWriterRepo)

	var ctx = context.WithValue(context.Background(), "X-Request-ID", "1234567890")
	ctx = context.WithValue(ctx, core.CurrentRequesterKeyStruct{}, core.RestRequester{
		ID: "1234567890",
	})

	t.Run("change password successfully", func(t *testing.T) {
		err := useCase.ExecChangePassword(ctx, "password", "new-password")

		assert.Nil(t, err)
	})
}

func TestResetPasswordUseCase(t *testing.T) {
	userReaderRepo := mockUserReaderRepository{}
	userWriterRepo := mockUserWriterRepository{}
	mailer := mockMailer{}
	HashAlgo := mockHashAlgo{}

	useCase := NewResetPasswordUseCase(mailer, HashAlgo, userReaderRepo, userWriterRepo)

	var ctx = context.WithValue(context.Background(), "X-Request-ID", "1234567890")
	t.Run("reset password successfully", func(t *testing.T) {
		err := useCase.ExecuteResetPassword(ctx, "existedemail@example.com")

		assert.Nil(t, err)
	})
}

func TestRefreshTokenUseCase(t *testing.T) {
	tokenProvider := mockTokenProvider{}

	useCase := NewRefreshTokenUseCase(tokenProvider, 1, 1)

	var ctx = context.WithValue(context.Background(), "X-Request-ID", "1234567890")
	ctx = context.WithValue(ctx, core.CurrentRequesterKeyStruct{}, core.RestRequester{
		ID: "1234567890",
	})

	t.Run("get refresh token successfully", func(t *testing.T) {
		accessToken, refreshToken, err := useCase.ExecRefreshToken(ctx)

		assert.Nil(t, err)
		assert.NotNil(t, accessToken)
		assert.NotNil(t, refreshToken)
	})
}
