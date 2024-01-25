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
	database "trekkstay/pkgs/db"
)

func TestFindHotelByID(t *testing.T) {
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

	repo := NewHotelRepoReader(*db)

	t.Run("should return hotels", func(t *testing.T) {
		hotel, err := repo.FindHotelByID(context.Background(), "25de6985-31b1-4f0d-82dd-25513bcb511b")

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

	connection := database.Connection{
		SSLMode:  database.Disable,
		Host:     dbConfig.DBHost,
		Port:     dbConfig.DBPort,
		Database: dbConfig.DBName,
		User:     dbConfig.DBUserName,
		Password: dbConfig.DBPassword,
	}

	db := database.InitDatabase(connection)

	repo := NewHotelRepoReader(*db)

	t.Run("should return hotels", func(t *testing.T) {
		paging, err := repo.FindHotels(context.Background(), entity.HotelFilterEntity{}, 1, 10)

		assert.Nil(t, err)
		assert.NotNil(t, paging)
	})
}

func TestFindRooms(t *testing.T) {
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
	repo := NewHotelRepoReader(*db)

	hotelID := "25de6985-31b1-4f0d-82dd-25513bcb511b"
	nonSmoking := false

	t.Run("should return rooms", func(t *testing.T) {
		rooms, err := repo.FindRooms(context.Background(), entity.RoomFilterEntity{
			HotelID:    &hotelID,
			NonSmoking: &nonSmoking,
		})

		fmt.Println("ROOMS: ", rooms)

		assert.Nil(t, err)
		assert.NotNil(t, rooms)
	})
}
