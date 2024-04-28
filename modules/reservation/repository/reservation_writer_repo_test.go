package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/modules/reservation/domain/entity"
	"trekkstay/pkgs/dbs/postgres"
)

func TestInsertReservation(t *testing.T) {
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
	repo := NewReservationWriterRepository(*db)

	t.Run("should insert reservation", func(t *testing.T) {
		err := repo.InsertReservation(context.Background(), &entity.ReservationEntity{
			RoomID:       "5e9847b7-72f4-4b04-97bf-afd8c08a5734",
			UserID:       "c2761241-2b31-47ff-8674-22570d7b495a",
			QRCodeURL:    "",
			Status:       "",
			CheckInDate:  "17-08-2021",
			CheckOutDate: "18-08-2021",
			Room: entity.RoomJSON{
				HotelID:       "093b02e4-8ca0-40f4-9133-c7ae5f733dd2",
				Type:          "Delix",
				OriginalPrice: 520000,
				BookingPrice:  520000,
				Images:        entity.MediaJSON{},
			},
		})

		assert.Nil(t, err)
	})
}

func TestUpdateReservation(t *testing.T) {
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
	repo := NewReservationWriterRepository(*db)

	t.Run("should update reservation", func(t *testing.T) {
		err := repo.UpdateReservationStatus(context.Background(),
			"1708277679316587562", "COMPLETED")

		assert.Nil(t, err)
	})
}
