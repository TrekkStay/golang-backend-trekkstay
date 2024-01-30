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
	"trekkstay/modules/hotel_room/domain/entity"
	"trekkstay/pkgs/dbs/postgres"
)

func TestCreateRoom(t *testing.T) {
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

	repo := NewHotelRoomRepoWriter(*db)

	t.Run("should insert room", func(t *testing.T) {
		err := repo.InsertHotelRoom(context.Background(), entity.HotelRoomEntity{
			HotelID:       "25de6985-31b1-4f0d-82dd-25513bcb511b",
			Type:          gofakeit.Name(),
			Quantity:      gofakeit.Number(1, 10),
			OriginalPrice: gofakeit.Number(1000000, 10000000),
			Videos:        entity.MediaJSON{},
			Images:        entity.MediaJSON{},
			Facilities: entity.HotelRoomFacilitiesJSON{
				RoomSize:    gofakeit.Number(10, 100),
				NumberOfBed: gofakeit.Number(1, 10),
				View:        gofakeit.RandomString([]string{"none", "city_view", "sea_view", "mountain_view"}),
				Balcony:     gofakeit.Bool(),
				BathTub:     gofakeit.Bool(),
				Kitchen:     gofakeit.Bool(),
				Television:  gofakeit.Bool(),
				Shower:      gofakeit.Bool(),
				NonSmoking:  gofakeit.Bool(),
				HairDryer:   gofakeit.Bool(),
				Sleeps: entity.SleepJSON{
					Adults:   gofakeit.Number(1, 4),
					Children: gofakeit.Number(1, 4),
				},
			},
		})

		assert.Nil(t, err)
	})
}

func TestUpdateRoom(t *testing.T) {
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

	repo := NewHotelRoomRepoWriter(*db)

	t.Run("should update room", func(t *testing.T) {
		err := repo.UpdateHotelRoom(context.Background(), entity.HotelRoomEntity{
			BaseEntity: core.BaseEntity{
				ID: "11ac7aa5-ad6d-4e44-a425-b480d4061f7d",
			},
			Type:          gofakeit.Name(),
			Description:   gofakeit.Paragraph(10, 20, 10, "\n"),
			Quantity:      gofakeit.Number(1, 10),
			OriginalPrice: gofakeit.Number(1000000, 10000000),
			Videos:        entity.MediaJSON{},
			Images:        entity.MediaJSON{},
			Facilities: entity.HotelRoomFacilitiesJSON{
				RoomSize:    gofakeit.Number(10, 100),
				NumberOfBed: gofakeit.Number(1, 10),
				View:        gofakeit.RandomString([]string{"none", "city_view", "sea_view", "mountain_view"}),
				Balcony:     gofakeit.Bool(),
				BathTub:     gofakeit.Bool(),
				Kitchen:     gofakeit.Bool(),
				Television:  gofakeit.Bool(),
				Shower:      gofakeit.Bool(),
			},
		})

		assert.Nil(t, err)
	})
}

func TestDeleteRoom(t *testing.T) {
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

	repo := NewHotelRoomRepoWriter(*db)

	t.Run("should delete room", func(t *testing.T) {
		err := repo.DeleteHotelRoom(context.Background(), "e962d40f-5b27-4c7f-9e66-0651e57804f6")
		assert.Nil(t, err)
	})
}
