package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
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
