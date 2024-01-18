package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/modules/user/domain/entity"
	database "trekkstay/pkgs/db"
)

func TestFindUserByCondition(t *testing.T) {
	err := os.Setenv("CONFIG_PATH", "../../../env/dev.env")
	if err != nil {
		return
	}

	dbConfig := config.LoadConfig(&models.DBConfig{}).(*models.DBConfig)

	connection := database.Connection{
		SSLMode:  database.Disable,
		Host:     dbConfig.DBHost,
		Port:     dbConfig.DBPort,
		Database: dbConfig.DBName,
		User:     dbConfig.DBUserName,
		Password: dbConfig.DBPassword,
	}

	db := database.InitDatabase(connection)

	userReaderRepo := NewUserReaderRepository(*db)
	userWriterRepo := NewUserWriterRepository(*db)

	_ = userWriterRepo.InsertUser(context.Background(), entity.UserEntity{
		FullName: "Test User",
		Email:    "testuser@example.com",
		Phone:    "1234567890",
		Status:   entity.ACTIVE.Value(),
		OTP:      "123456",
		Password: "password",
		Salt:     "salt",
	})

	t.Run("find user by email", func(t *testing.T) {
		condition := map[string]interface{}{
			"email": "testuser@example.com",
		}

		_, err := userReaderRepo.FindUserByCondition(context.Background(), condition)

		assert.Nil(t, err)
	})

	t.Run("find user by phone", func(t *testing.T) {
		condition := map[string]interface{}{
			"phone": "7350510091",
		}

		_, err := userReaderRepo.FindUserByCondition(context.Background(), condition)

		assert.Nil(t, err)
	})
}
