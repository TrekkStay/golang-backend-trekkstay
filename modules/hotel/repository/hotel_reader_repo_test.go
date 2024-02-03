package repository

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/modules/hotel/domain/entity"
	"trekkstay/pkgs/dbs/postgres"
)

func TestFindHotelByID(t *testing.T) {
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

	repo := NewHotelReaderRepository(*db)

	t.Run("should return hotels", func(t *testing.T) {
		hotel, err := repo.FindHotelByCondition(context.Background(), map[string]interface{}{
			"id": "25de6985-31b1-4f0d-82dd-25513bcb511b",
		})

		assert.Nil(t, err)
		assert.NotNil(t, hotel)
	})
}

func TestPagingHotel(t *testing.T) {
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

	repo := NewHotelReaderRepository(*db)

	t.Run("should return hotels", func(t *testing.T) {
		paging, err := repo.FindHotels(context.Background(), entity.HotelFilterEntity{}, 1, 10)

		fmt.Print(paging)

		assert.Nil(t, err)
		assert.NotNil(t, paging)
	})
}
