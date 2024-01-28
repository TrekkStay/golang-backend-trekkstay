package repository

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/modules/user/domain/entity"
	"trekkstay/pkgs/dbs/postgres"
)

func TestInsertUser(t *testing.T) {
	err := os.Setenv("CONFIG_PATH", "../../../env/dev.env")
	if err != nil {
		return
	}

	dbConfig := config.LoadConfig(&models.DBConfig{}).(*models.DBConfig)

	connection := postgres.Connection{
		SSLMode:  postgres.Disable,
		Host:     dbConfig.DBHost,
		Port:     dbConfig.DBPort,
		Database: dbConfig.DBName,
		User:     dbConfig.DBUserName,
		Password: dbConfig.DBPassword,
	}

	db := postgres.InitDatabase(connection)

	userRepo := NewUserWriterRepository(*db)

	t.Run("should insert user", func(t *testing.T) {
		err := userRepo.InsertUser(context.Background(), entity.UserEntity{
			FullName: gofakeit.Name(),
			Email:    gofakeit.Email(),
			Phone:    gofakeit.Phone(),
			Status:   gofakeit.RandomString([]string{entity.ACTIVE.Value(), entity.UNVERIFIED.Value(), entity.BLOCKED.Value()}),
			OTP:      "123456",
			Password: gofakeit.Password(true, true, true, false, false, 10),
		})

		assert.Nil(t, err)
	})

	t.Run("should not insert user", func(t *testing.T) {
		_ = userRepo.InsertUser(context.Background(), entity.UserEntity{
			FullName: gofakeit.Name(),
			Email:    "testuser@example.com",
			Phone:    gofakeit.Phone(),
			Status:   gofakeit.RandomString([]string{entity.ACTIVE.Value(), entity.UNVERIFIED.Value(), entity.BLOCKED.Value()}),
			OTP:      "123456",
			Password: gofakeit.Password(true, true, true, false, false, 10),
		})

		// insert duplicate user
		err := userRepo.InsertUser(context.Background(), entity.UserEntity{
			FullName: gofakeit.Name(),
			Email:    "testuser@example.com",
			Phone:    gofakeit.Phone(),
			Status:   gofakeit.RandomString([]string{entity.ACTIVE.Value(), entity.UNVERIFIED.Value(), entity.BLOCKED.Value()}),
			OTP:      "123456",
			Password: gofakeit.Password(true, true, true, false, false, 10),
		})

		assert.NotNil(t, err)
	})
}

func TestUpdateUser(t *testing.T) {
	err := os.Setenv("CONFIG_PATH", "../../../env/dev.env")
	if err != nil {
		return
	}

	dbConfig := config.LoadConfig(&models.DBConfig{}).(*models.DBConfig)

	connection := postgres.Connection{
		SSLMode:  postgres.Disable,
		Host:     dbConfig.DBHost,
		Port:     dbConfig.DBPort,
		Database: dbConfig.DBName,
		User:     dbConfig.DBUserName,
		Password: dbConfig.DBPassword,
	}

	db := postgres.InitDatabase(connection)

	userRepo := NewUserWriterRepository(*db)

	t.Run("should update user", func(t *testing.T) {
		_ = userRepo.InsertUser(context.Background(), entity.UserEntity{
			FullName: gofakeit.Name(),
			Email:    "testuser@example.com",
			Phone:    gofakeit.Phone(),
			Status:   gofakeit.RandomString([]string{entity.ACTIVE.Value(), entity.UNVERIFIED.Value(), entity.BLOCKED.Value()}),
			OTP:      "123456",
			Password: gofakeit.Password(true, true, true, false, false, 10),
		})

		err := userRepo.UpdateUser(context.Background(), entity.UserEntity{
			FullName: gofakeit.Name(),
			Email:    "testuser@example.com",
			Phone:    gofakeit.Phone(),
			Status:   gofakeit.RandomString([]string{entity.ACTIVE.Value(), entity.UNVERIFIED.Value(), entity.BLOCKED.Value()}),
			OTP:      "123456",
			Password: gofakeit.Password(true, true, true, false, false, 10),
		})

		assert.Nil(t, err)
	})
}

func TestDeleteUser(t *testing.T) {
	err := os.Setenv("CONFIG_PATH", "../../../env/dev.env")
	if err != nil {
		return
	}

	dbConfig := config.LoadConfig(&models.DBConfig{}).(*models.DBConfig)

	connection := postgres.Connection{
		SSLMode:  postgres.Disable,
		Host:     dbConfig.DBHost,
		Port:     dbConfig.DBPort,
		Database: dbConfig.DBName,
		User:     dbConfig.DBUserName,
		Password: dbConfig.DBPassword,
	}

	db := postgres.InitDatabase(connection)

	userRepo := NewUserWriterRepository(*db)

	t.Run("should delete user", func(t *testing.T) {
		_ = userRepo.InsertUser(context.Background(), entity.UserEntity{
			BaseEntity: baseentity.BaseEntity{
				ID: "3a1d7b7f-1eeb-41d2-9d9e-4cce3785ee01",
			},
			FullName: gofakeit.Name(),
			Email:    "testuser@example.com",
			Phone:    gofakeit.Phone(),
			Status:   gofakeit.RandomString([]string{entity.ACTIVE.Value(), entity.UNVERIFIED.Value(), entity.BLOCKED.Value()}),
			OTP:      "123456",
			Password: gofakeit.Password(true, true, true, false, false, 10),
		})

		err := userRepo.DeleteUser(context.Background(), "3a1d7b7f-1eeb-41d2-9d9e-4cce3785ee01")

		assert.Nil(t, err)
	})
}
