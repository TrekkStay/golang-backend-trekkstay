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

func TestIRFilterHotel(t *testing.T) {
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

	useCase := usecase.NewFilterHotelUseCase(hotelReaderRepo)

	var ctx = context.WithValue(context.Background(), "X-Request-ID", "1234567890")
	ctx = context.WithValue(ctx, core.CurrentRequesterKeyStruct{}, core.RestRequester{
		ID:   "19bd8d4f-ed6c-4b65-a181-dd38f6b809dc",
		Role: "OWNER",
	})

	name := "Anrizon"
	_, err = useCase.FilterHotel(ctx, entity.HotelFilterEntity{
		Name: &name,
	}, 1, 10)

	assert.Nil(t, err)
}
