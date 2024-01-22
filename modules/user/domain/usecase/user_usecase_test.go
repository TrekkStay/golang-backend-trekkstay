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

	ctx := context.WithValue(context.Background(), "X-Request-ID", "1234567890")

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

	ctx := context.WithValue(context.Background(), "X-Request-ID", "1234567890")

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

	ctx := context.WithValue(context.Background(), "X-Request-ID", "1234567890")
	ctx = context.WithValue(ctx, core.CurrentRequesterKey, core.RestRequester{
		Id: "1234567890",
	})

	t.Run("change password successfully", func(t *testing.T) {
		err := useCase.ExecChangePassword(ctx, "password", "new-password")

		assert.Nil(t, err)
	})
}

func TestForgotPasswordUseCase(t *testing.T) {
	userReaderRepo := mockUserReaderRepository{}
	userWriterRepo := mockUserWriterRepository{}
	mailer := mockMailer{}

	useCase := NewForgotPasswordUseCase(mailer, userReaderRepo, userWriterRepo)

	ctx := context.WithValue(context.Background(), "X-Request-ID", "1234567890")

	t.Run("forgot password successfully", func(t *testing.T) {
		err := useCase.ExecuteForgotPassword(ctx, "existedemail@example.com")

		assert.Nil(t, err)
	})
}
