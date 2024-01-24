package api

import (
	"time"
	"trekkstay/api/routes"
	"trekkstay/config"
	"trekkstay/config/models"
	regionHandler "trekkstay/modules/region/api/handler"
	userHandler "trekkstay/modules/user/api/handler"
	userUseCase "trekkstay/modules/user/domain/usecase"
	userRepo "trekkstay/modules/user/repository"

	regionUseCase "trekkstay/modules/region/domain/usecase"
	regionRepo "trekkstay/modules/region/repository"

	database "trekkstay/pkgs/db"
	"trekkstay/pkgs/jwt"
	"trekkstay/pkgs/mail"
	"trekkstay/pkgs/transport/http/server"
	"trekkstay/utils"
)

func NewServer() (*server.HTTPServer, error) {
	appConfig := config.LoadConfig(&models.AppConfig{}).(*models.AppConfig)
	jwtConfig := config.LoadConfig(&models.JWTConfig{}).(*models.JWTConfig)
	dbConfig := config.LoadConfig(&models.DBConfig{}).(*models.DBConfig)
	mailConfig := config.LoadConfig(&models.MailConfig{}).(*models.MailConfig)

	connection := database.Connection{
		SSLMode:               database.Disable,
		Host:                  dbConfig.DBHost,
		Port:                  dbConfig.DBPort,
		Database:              dbConfig.DBName,
		User:                  dbConfig.DBUserName,
		Password:              dbConfig.DBPassword,
		MaxIdleConnections:    dbConfig.MaxIdleConnections,
		MaxOpenConnections:    dbConfig.MaxOpenConnections,
		ConnectionMaxIdleTime: time.Duration(dbConfig.ConnectionMaxIdleTime),
		ConnectionMaxLifeTime: time.Duration(dbConfig.ConnectionMaxLifeTime),
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
	mailer := mail.NewMailer(mailConfig)

	// User Repository
	userRepoReader := userRepo.NewUserReaderRepository(*db)
	userRepoWriter := userRepo.NewUserWriterRepository(*db)

	// Region Repository
	regionRepoReader := regionRepo.NewRegionReaderRepository(*db)

	srv := &routes.RouteHandler{
		UserHandler: userHandler.NewUserHandler(requestValidator,
			userUseCase.NewCreateUserUseCase(hashAlgo, userRepoReader, userRepoWriter),
			userUseCase.NewLoginUserUseCase(jwtToken, jwtConfig.AccessTokenExpiry,
				jwtConfig.RefreshTokenExpiry, hashAlgo, userRepoReader),
			userUseCase.NewChangePasswordUseCase(hashAlgo, userRepoReader, userRepoWriter),
			userUseCase.NewForgotPasswordUseCase(mailer, hashAlgo, userRepoReader, userRepoWriter),
		),
		RegionHandler: regionHandler.NewRegionHandler(
			regionUseCase.NewListProvinceUseCase(regionRepoReader),
			regionUseCase.NewListDistrictUseCase(regionRepoReader),
			regionUseCase.NewListWardUseCase(regionRepoReader),
		),
	}

	s.AddRoutes(srv.InitRoutes())
	s.AddGroupRoutes(srv.InitGroupRoutes())

	return s, nil
}
