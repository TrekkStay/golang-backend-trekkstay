package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/modules/hotel_room/domain/entity"
	"trekkstay/pkgs/dbs/postgres"
)

func TestFindRooms(t *testing.T) {
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
	repo := NewHotelRoomReaderRepository(*db)

	hotelID := "1a21fdb9-b9e1-4d64-bc39-26930911ce06"
	nonSmoking := false

	t.Run("should return rooms", func(t *testing.T) {
		rooms, err := repo.FindHotelRooms(context.Background(), entity.HotelRoomFilterEntity{
			HotelID:    &hotelID,
			NonSmoking: &nonSmoking,
		})

		assert.Nil(t, err)
		assert.NotNil(t, rooms)
	})
}
