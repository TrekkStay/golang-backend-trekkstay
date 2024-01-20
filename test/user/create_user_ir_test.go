package user

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"sync"
	"testing"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/modules/user/domain/entity"
	"trekkstay/modules/user/domain/usecase"
	"trekkstay/modules/user/repository"
	database "trekkstay/pkgs/db"
	"trekkstay/utils"
)

func TestIRCreateUser(t *testing.T) {
	err := os.Setenv("CONFIG_PATH", "../../env/dev.env")
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

	userReaderRepo := repository.NewUserReaderRepository(*db)
	userWriterRepo := repository.NewUserWriterRepository(*db)
	hashAlgo := utils.NewHashAlgo()

	useCase := usecase.NewCreateUserUseCase(hashAlgo, userReaderRepo, userWriterRepo)

	ctx := context.WithValue(context.Background(), "X-Request-ID", "1234567890")

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			err := useCase.ExecCreateUser(ctx, entity.UserEntity{
				FullName: gofakeit.Name(),
				Email:    gofakeit.Email(),
				Phone:    gofakeit.Phone(),
				Status: gofakeit.RandomString([]string{
					entity.ACTIVE.Value(),
					entity.INACTIVE.Value(),
					entity.BLOCKED.Value(),
				}),
				OTP:      strconv.Itoa(gofakeit.RandomInt([]int{100000, 999999})),
				Password: gofakeit.Password(true, true, true, false, false, 10),
			})

			assert.Nil(t, err)
		}()

		wg.Wait()
	}
}
