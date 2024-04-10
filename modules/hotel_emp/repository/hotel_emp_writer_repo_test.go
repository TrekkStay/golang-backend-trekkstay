package repository

import (
	"context"
	"os"
	"testing"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/core"
	"trekkstay/modules/hotel_emp/domain/entity"
	"trekkstay/pkgs/dbs/postgres"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestCreateHotelEmp(t *testing.T) {
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

	repo := NewHotelEmpWriterRepository(*db)

	t.Run("should insert hotel employee", func(t *testing.T) {
		err := repo.InsertHotelEmp(context.Background(), entity.HotelEmpEntity{
			FullName:   gofakeit.Name(),
			Email:      gofakeit.Email(),
			Phone:      gofakeit.Phone(),
			Contract:   gofakeit.RandomString([]string{"full-time, part-time"}),
			BaseSalary: gofakeit.Number(1000000, 10000000),
			Role:       "owner",
			Status:     "active",
			OTP:        "123123",
			Password:   "12312312",
		})

		assert.Nil(t, err)
	})
}

func TestUpdateHotelEmp(t *testing.T) {
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

	repo := NewHotelEmpWriterRepository(*db)

	t.Run("should update hotel employee", func(t *testing.T) {
		err := repo.UpdateHotelEmp(context.Background(), entity.HotelEmpEntity{
			BaseEntity: core.BaseEntity{
				ID: "d4805d31-4c90-4fd6-8a1a-c7c96b63e54e",
			},
			FullName:   gofakeit.Name(),
			Email:      gofakeit.Email(),
			Phone:      gofakeit.Phone(),
			Contract:   gofakeit.RandomString([]string{"full-time, part-time"}),
			BaseSalary: gofakeit.Number(1000000, 10000000),
			Role:       "owner",
			Status:     "active",
			OTP:        "123123",
			Password:   "12312312",
		})

		assert.Nil(t, err)
	})
}

func TestDeleteHotelEmp(t *testing.T) {
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

	repo := NewHotelEmpWriterRepository(*db)

	t.Run("should delete hotel employee", func(t *testing.T) {
		err := repo.DeleteHotelEmp(context.Background(), "d4805d31-4c90-4fd6-8a1a-c7c96b63e54e")
		assert.Nil(t, err)
	})
}
