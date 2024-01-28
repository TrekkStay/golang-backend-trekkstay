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

	repo := NewHotelRepoWriter(*db)

	t.Run("should insert hotel employee", func(t *testing.T) {
		err := repo.InsertHotelEmployee(context.Background(), entity.HotelEmployeeEntity{
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
			OwnerID:       "d4805d31-4c90-4fd6-8a1a-c7c96b63e54e",
			Email:         gofakeit.Email(),
			Phone:         gofakeit.Phone(),
			ProvinceCode:  "01",
			DistrictCode:  "001",
			WardCode:      "00001",
			AddressDetail: "00001",
			Description:   gofakeit.Sentence(10),
			Status:        "active",
			HotelFacility: entity.HotelFacilityEntity{
				MotorBikeRental: true,
				LaundryService:  true,
				FreeWifi:        true,
			},
			Coordinates: entity.CoordinatesEntity{
				Lat: 10.123123,
				Lng: 10.123123,
			},
			Videos: entity.MediaObject{
				URL: []string{"https://www.youtube.com/watch?v=1"},
			},
			Images: entity.MediaObject{
				URL: []string{"https://www.youtube.com/watch?v=1"},
			},
		})

		assert.Nil(t, err)
	})
}

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

	repo := NewHotelRepoWriter(*db)

	t.Run("should insert room", func(t *testing.T) {
		err := repo.InsertRoom(context.Background(), entity.RoomEntity{
			HotelID:     "25de6985-31b1-4f0d-82dd-25513bcb511b",
			Type:        gofakeit.Name(),
			Quantity:    gofakeit.Number(1, 10),
			OriginPrice: gofakeit.Number(1000000, 10000000),
			Videos:      entity.MediaObject{},
			Images:      entity.MediaObject{},
			RoomFacility: entity.RoomFacilityEntity{
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
				Sleeps: entity.SleepsEntity{
					Adults:   gofakeit.Number(1, 4),
					Children: gofakeit.Number(1, 4),
				},
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
			HotelFacility: entity.HotelFacilityEntity{
				ParkingArea:     true,
				MotorBikeRental: true,
				LaundryService:  true,
				FreeWifi:        true,
			},
			Coordinates: entity.CoordinatesEntity{
				Lat: 10.123123,
				Lng: 10.123123,
			},
			Videos: entity.MediaObject{
				URL: []string{"https://www.youtube.com/watch?v=1"},
			},
			Images: entity.MediaObject{
				URL: []string{"https://www.youtube.com/watch?v=1"},
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

	repo := NewHotelRepoWriter(*db)

	t.Run("should update room", func(t *testing.T) {
		err := repo.UpdateRoom(context.Background(), entity.RoomEntity{
			BaseEntity: core.BaseEntity{
				ID: "e962d40f-5b27-4c7f-9e66-0651e57804f6",
			},
			Type:        gofakeit.Name(),
			Description: gofakeit.Paragraph(10, 20, 10, "\n"),
			Quantity:    gofakeit.Number(1, 10),
			OriginPrice: gofakeit.Number(1000000, 10000000),
			Videos:      entity.MediaObject{},
			Images:      entity.MediaObject{},
			RoomFacility: entity.RoomFacilityEntity{
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

	repo := NewHotelRepoWriter(*db)

	t.Run("should update hotel employee", func(t *testing.T) {
		err := repo.UpdateHotelEmployee(context.Background(), entity.HotelEmployeeEntity{
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

	repo := NewHotelRepoWriter(*db)

	t.Run("should delete room", func(t *testing.T) {
		err := repo.DeleteRoom(context.Background(), "e962d40f-5b27-4c7f-9e66-0651e57804f6")
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

	repo := NewHotelRepoWriter(*db)

	t.Run("should delete hotel employee", func(t *testing.T) {
		err := repo.DeleteHotelEmployee(context.Background(), "d4805d31-4c90-4fd6-8a1a-c7c96b63e54e")
		assert.Nil(t, err)
	})
}
