package api

import (
	"time"
	"trekkstay/api/routes"
	"trekkstay/config"
	"trekkstay/config/models"
	userHandler "trekkstay/modules/user/api/handler"
	userUseCase "trekkstay/modules/user/domain/usecase"
	userRepo "trekkstay/modules/user/repository"
	database "trekkstay/pkgs/db"
	"trekkstay/pkgs/jwt"
	"trekkstay/pkgs/transport/http/server"
	"trekkstay/utils"
)

func NewServer() (*server.HTTPServer, error) {
	appConfig := config.LoadConfig(&models.AppConfig{}).(*models.AppConfig)
	jwtConfig := config.LoadConfig(&models.JWTConfig{}).(*models.JWTConfig)
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

	s := server.NewHTTPServer(
		server.AddName(appConfig.ServiceName),
		server.AddPort(appConfig.ServicePort),
		server.SetGracefulShutdownTimeout(time.Duration(appConfig.ServiceTimeout)),
	)

	requestValidator := utils.NewValidator()
	hashAlgo := utils.NewHashAlgo()
	jwtToken := jwt.NewJWT(jwtConfig.JWTSecretKey)

	// User Repository
	userRepoReader := userRepo.NewUserReaderRepository(*db)
	userRepoWriter := userRepo.NewUserWriterRepository(*db)

	srv := &routes.RouteHandler{
		UserHandler: userHandler.NewUserHandler(requestValidator,
			userUseCase.NewCreateUserUseCase(hashAlgo, userRepoReader, userRepoWriter),
			userUseCase.NewLoginUserUseCase(jwtToken, jwtConfig.AccessTokenExpiry,
				jwtConfig.RefreshTokenExpiry, hashAlgo, userRepoReader),
		),
	}

	s.AddRoutes(srv.InitRoutes())
	s.AddGroupRoutes(srv.InitGroupRoutes())

	return s, nil
}
