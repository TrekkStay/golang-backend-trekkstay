package user

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/core"
	"trekkstay/modules/user/domain/entity"
	"trekkstay/modules/user/domain/usecase"
	"trekkstay/modules/user/repository"
	"trekkstay/pkgs/dbs/postgres"
)

func TestIRUpdateUser(t *testing.T) {
	err := os.Setenv("CONFIG_PATH", "../../env/dev.env")
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

	userReaderRepo := repository.NewUserReaderRepository(*db)
	userWriterRepo := repository.NewUserWriterRepository(*db)

	useCase := usecase.NewUpdateUserUseCase(userReaderRepo, userWriterRepo)

	var ctx = context.WithValue(context.Background(), "X-Request-ID", "1234567890")
	ctx = context.WithValue(ctx, core.CurrentRequesterKeyStruct{}, core.RestRequester{
		ID: "1cd51b28-aa1f-4105-ae74-0abfa050f621",
	})

	err = useCase.ExecUpdateUser(ctx, entity.UserEntity{
		FullName: gofakeit.Name(),
		Email:    gofakeit.Email(),
	})

	assert.Nil(t, err)
}
