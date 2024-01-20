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
	"trekkstay/modules/user/domain/entity"
	"trekkstay/modules/user/domain/usecase"
	"trekkstay/modules/user/repository"
	database "trekkstay/pkgs/db"
	"trekkstay/pkgs/jwt"
	"trekkstay/utils"
)

func TestIRLoginUser(t *testing.T) {
	err := os.Setenv("CONFIG_PATH", "../../env/dev.env")
	if err != nil {
		return
	}

	dbConfig := config.LoadConfig(&models.DBConfig{}).(*models.DBConfig)
	jwtConfig := config.LoadConfig(&models.JWTConfig{}).(*models.JWTConfig)

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
	jwtToken := jwt.NewJWT(jwtConfig.JWTSecretKey)

	loginUserUseCase := usecase.NewLoginUserUseCase(jwtToken, jwtConfig.AccessTokenExpiry,
		jwtConfig.RefreshTokenExpiry, hashAlgo, userReaderRepo)

	createUserUseCase := usecase.NewCreateUserUseCase(hashAlgo, userReaderRepo, userWriterRepo)

	ctx := context.WithValue(context.Background(), "X-Request-ID", "1234567890")

	_ = createUserUseCase.ExecCreateUser(ctx, entity.UserEntity{
		FullName: gofakeit.Name(),
		Email:    "testlogin@example.com",
		Phone:    gofakeit.Phone(),
		Status: gofakeit.RandomString([]string{
			entity.ACTIVE.Value(),
			entity.INACTIVE.Value(),
			entity.BLOCKED.Value(),
		}),
		OTP:      strconv.Itoa(gofakeit.RandomInt([]int{100000, 999999})),
		Password: "test_login_password",
	})

	_, err = loginUserUseCase.ExecLoginUser(ctx, entity.UserEntity{
		Email:    "testlogin@example.com",
		Password: "test_login_password",
	})

	assert.Nil(t, err)
}
