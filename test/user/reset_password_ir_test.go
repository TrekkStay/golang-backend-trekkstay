package user

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/modules/user/domain/usecase"
	"trekkstay/modules/user/repository"
	"trekkstay/pkgs/dbs/postgres"
	"trekkstay/pkgs/mail"
	"trekkstay/utils"
)

func TestIRResetPassword(t *testing.T) {
	err := os.Setenv("CONFIG_PATH", "../../env/dev.env")
	if err != nil {
		return
	}

	dbConfig := config.LoadConfig(&models.DBConfig{}).(*models.DBConfig)
	mailConfig := config.LoadConfig(&models.MailConfig{}).(*models.MailConfig)

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
	mailer := mail.NewMailer(mailConfig)
	hashAlgo := utils.NewHashAlgo()

	forgotPasswordUseCase := usecase.NewResetPasswordUseCase(mailer, hashAlgo, userReaderRepo, userWriterRepo)

	var ctx = context.WithValue(context.Background(), "X-Request-ID", "1234567890")
	err = forgotPasswordUseCase.ExecuteResetPassword(ctx, "thanhanphan17@gmail.com")

	assert.Nil(t, err)
}
