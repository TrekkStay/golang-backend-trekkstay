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

	useCase := NewCreateUserUseCase(userReaderRepo, userWriterRepo)

	t.Run("create user successfully", func(t *testing.T) {
		err := useCase.ExecCreateUser(context.Background(), entity.UserEntity{
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
		err := useCase.ExecCreateUser(context.Background(), entity.UserEntity{
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
