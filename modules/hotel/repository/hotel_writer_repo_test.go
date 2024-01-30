package repository

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/core"
	"trekkstay/modules/hotel/domain/entity"
	"trekkstay/pkgs/dbs/postgres"
)

func TestCreateHotel(t *testing.T) {
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

	repo := NewHotelRepoWriter(*db)

	t.Run("should insert hotel", func(t *testing.T) {
		err := repo.InsertHotel(context.Background(), entity.HotelEntity{
			Name:          gofakeit.Name(),
			OwnerID:       "b4f480bd-f3d8-4a4c-a436-0b78b20b95e0",
			Email:         gofakeit.Email(),
			Phone:         gofakeit.Phone(),
			ProvinceCode:  "01",
			DistrictCode:  "001",
			WardCode:      "00001",
			AddressDetail: "00001",
			Description:   gofakeit.Sentence(10),
			Status:        "active",
			Facilities: entity.HotelFacilitiesJSON{
				MotorBikeRental: true,
				LaundryService:  true,
				FreeWifi:        true,
			},
			Coordinates: entity.CoordinatesJSON{
				Lat: 10.123123,
				Lng: 10.123123,
			},
			Videos: entity.MediaJSON{
				URL: []string{"https://www.youtube.com/watch?v=1"},
			},
			Images: entity.MediaJSON{
				URL: []string{"https://www.youtube.com/watch?v=1"},
			},
		})

		assert.Nil(t, err)
	})
}

func TestUpdateHotel(t *testing.T) {
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

	repo := NewHotelRepoWriter(*db)

	t.Run("should update hotel", func(t *testing.T) {
		err := repo.UpdateHotel(context.Background(), entity.HotelEntity{
			BaseEntity: core.BaseEntity{
				ID: "25de6985-31b1-4f0d-82dd-25513bcb511b",
			},
			Name:          gofakeit.Name(),
			OwnerID:       "d4805d31-4c90-4fd6-8a1a-c7c96b63e54e",
			Email:         gofakeit.Email(),
			Phone:         gofakeit.Phone(),
			ProvinceCode:  "01",
			DistrictCode:  "001",
			WardCode:      "00001",
			AddressDetail: "00001",
			Description:   gofakeit.Sentence(10),
			Status:        "active",
			Facilities: entity.HotelFacilitiesJSON{
				ParkingArea:     true,
				MotorBikeRental: true,
				LaundryService:  true,
				FreeWifi:        true,
			},
			Coordinates: entity.CoordinatesJSON{
				Lat: 10.123123,
				Lng: 10.123123,
			},
			Videos: entity.MediaJSON{
				URL: []string{"https://www.youtube.com/watch?v=1"},
			},
			Images: entity.MediaJSON{
				URL: []string{"https://www.youtube.com/watch?v=1"},
			},
		})

		assert.Nil(t, err)
	})
}

func TestDeleteHotel(t *testing.T) {
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

	repo := NewHotelRepoWriter(*db)

	t.Run("should delete hotel", func(t *testing.T) {
		err := repo.DeleteHotel(context.Background(), "25de6985-31b1-4f0d-82dd-25513bcb511b")
		assert.Nil(t, err)
	})
}
