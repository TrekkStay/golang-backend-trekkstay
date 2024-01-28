package user

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/core"
	"trekkstay/modules/user/domain/entity"
	"trekkstay/modules/user/domain/usecase"
	"trekkstay/modules/user/repository"
	"trekkstay/pkgs/dbs/postgres"
	"trekkstay/utils"
)

func TestIRChangePassword(t *testing.T) {
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
	hashAlgo := utils.NewHashAlgo()

	createUserUseCase := usecase.NewCreateUserUseCase(hashAlgo, userReaderRepo, userWriterRepo)
	changePasswordUseCase := usecase.NewChangePasswordUseCase(hashAlgo, userReaderRepo, userWriterRepo)

	ctx := context.WithValue(context.Background(), "X-Request-ID", "1234567890")

	_ = createUserUseCase.ExecCreateUser(ctx, entity.UserEntity{
		BaseEntity: core.BaseEntity{
			ID: "151d3f25-7c4e-4c9a-a3b8-55356ebcfbf56",
		},
		FullName: gofakeit.Name(),
		Email:    "testchangepassword@example.com",
		Phone:    gofakeit.Phone(),
		Status: gofakeit.RandomString([]string{
			entity.ACTIVE.Value(),
			entity.INACTIVE.Value(),
			entity.BLOCKED.Value(),
		}),
		OTP:      strconv.Itoa(gofakeit.RandomInt([]int{100000, 999999})),
		Password: "1234567890",
	})

	ctx = context.WithValue(ctx, core.CurrentRequesterKeyStruct{}, core.RestRequester{
		ID: "151d3f25-7c4e-4c9a-a3b8-55356ebcfbf56",
	})

	err = changePasswordUseCase.ExecChangePassword(ctx, "1234567890", "123456789012")

	assert.Nil(t, err)
}
