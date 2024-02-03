package hotel_emp

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/core"
	"trekkstay/modules/hotel_emp/domain/entity"
	"trekkstay/modules/hotel_emp/domain/usecase"
	"trekkstay/modules/hotel_emp/repository"
	"trekkstay/pkgs/dbs/postgres"
	"trekkstay/pkgs/mail"
	"trekkstay/utils"
)

func TestIRCreateHotelEmp(t *testing.T) {
	err := os.Setenv("CONFIG_PATH", "../../env/dev.env")
	if err != nil {
		return
	}

	dbConfig := config.LoadConfig(&models.DBConfig{}).(*models.DBConfig)
	mailConfig := config.LoadConfig(&models.MailConfig{}).(*models.MailConfig)

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

	hotelEmpReaderRepo := repository.NewHotelEmpRepoReader(*db)
	hotelEmpWriterRepo := repository.NewHotelEmpRepoWriter(*db)
	hashAlgo := utils.NewHashAlgo()
	mailer := mail.NewMailer(mailConfig)

	useCase := usecase.NewCreateHotelEmpUseCase(mailer, hashAlgo, hotelEmpReaderRepo, hotelEmpWriterRepo)

	var ctx = context.WithValue(context.Background(), "X-Request-ID", "1234567890")
	ctx = context.WithValue(ctx, core.CurrentRequesterKeyStruct{}, core.RestRequester{
		ID: "05ccec7e-31bb-40db-b814-bb4e90e855bf",
	})

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			err := useCase.ExecuteCreateHotelEmp(ctx, entity.HotelEmpEntity{
				FullName: gofakeit.Name(),
				Email:    gofakeit.Email(),
				Phone:    gofakeit.Phone(),
				Status: gofakeit.RandomString([]string{
					entity.ACTIVE.Value(),
					entity.BLOCKED.Value(),
					entity.UNVERIFIED.Value(),
				}),
				OTP:      strconv.Itoa(gofakeit.RandomInt([]int{100000, 999999})),
				Password: gofakeit.Password(true, true, true, false, false, 10),
			})

			assert.Nil(t, err)
		}()

		wg.Wait()
	}
}
