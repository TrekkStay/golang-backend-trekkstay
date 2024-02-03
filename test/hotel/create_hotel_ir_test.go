package hotel

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/core"
	"trekkstay/modules/hotel/domain/entity"
	"trekkstay/modules/hotel/domain/usecase"
	"trekkstay/modules/hotel/repository"
	"trekkstay/pkgs/dbs/postgres"
)

func TestIRCreateHotel(t *testing.T) {
	err := os.Setenv("CONFIG_PATH", "../../env/dev.env")
	if err != nil {
		return
	}

	dbConfig := config.LoadConfig(&models.DBConfig{}).(*models.DBConfig)

	connection := postgres.Connection{
		SSLMode:               postgres.Disable,
		Host:                  dbConfig.DBHost,
		Port:                  dbConfig.DBPort,
		Database:              dbConfig.DBName,
		User:                  dbConfig.DBUserName,
		Password:              dbConfig.DBPassword,
		MaxIdleConnections:    dbConfig.MaxIdleConnections,
		MaxOpenConnections:    dbConfig.MaxOpenConnections,
		ConnectionMaxIdleTime: time.Duration(dbConfig.ConnectionMaxIdleTime),
		ConnectionMaxLifeTime: time.Duration(dbConfig.ConnectionMaxLifeTime),
		ConnectionTimeout:     time.Duration(dbConfig.ConnectionTimeout),
	}

	db := postgres.InitDatabase(connection)

	hotelReaderRepo := repository.NewHotelReaderRepository(*db)
	hotelWriterRepo := repository.NewHotelWriterRepository(*db)

	useCase := usecase.NewCreateHotelUseCase(hotelReaderRepo, hotelWriterRepo)

	var ctx = context.WithValue(context.Background(), "X-Request-ID", "1234567890")
	ctx = context.WithValue(ctx, core.CurrentRequesterKeyStruct{}, core.RestRequester{
		ID:   "19bd8d4f-ed6c-4b65-a181-dd38f6b809dc",
		Role: "OWNER",
	})

	err = useCase.ExecuteCreateHotel(ctx, entity.HotelEntity{
		Name:          "Anrizon Hotel",
		Email:         "anrizon@anrizon.com",
		Phone:         "023456789",
		CheckInTime:   "12:00",
		CheckOutTime:  "12:00",
		ProvinceCode:  "01",
		DistrictCode:  "001",
		WardCode:      "00001",
		AddressDetail: "So 4, Pham Van Dong",
		Description:   "Gan Duong Tran Phu",
		Facilities: entity.HotelFacilitiesJSON{
			FreeWifi: true,
		},
		Coordinates: entity.CoordinatesJSON{},
		Videos:      entity.MediaJSON{},
		Images:      entity.MediaJSON{},
	})

	assert.Nil(t, err)
}
