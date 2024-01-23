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
	database "trekkstay/pkgs/db"
	"trekkstay/pkgs/mail"
)

func TestIRForgotPassword(t *testing.T) {
	err := os.Setenv("CONFIG_PATH", "../../env/dev.env")
	if err != nil {
		return
	}

	dbConfig := config.LoadConfig(&models.DBConfig{}).(*models.DBConfig)
	mailConfig := config.LoadConfig(&models.MailConfig{}).(*models.MailConfig)

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
	mailer := mail.NewMailer(mailConfig)

	forgotPasswordUseCase := usecase.NewForgotPasswordUseCase(mailer, userReaderRepo, userWriterRepo)

	ctx := context.WithValue(context.Background(), "X-Request-ID", "1234567890")

	err = forgotPasswordUseCase.ExecuteForgotPassword(ctx, "thanhanphan17@gmail.com")

	assert.Nil(t, err)
}