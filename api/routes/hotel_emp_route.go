package routes

import (
	"github.com/go-playground/validator/v10"
	"trekkstay/api/middlewares"
	"trekkstay/config"
	"trekkstay/config/models"
	hotelEmpHandler "trekkstay/modules/hotel_emp/api/handler"
	"trekkstay/modules/hotel_emp/domain/usecase"
	"trekkstay/modules/hotel_emp/repository"
	database "trekkstay/pkgs/dbs/postgres"
	"trekkstay/pkgs/jwt"
	"trekkstay/pkgs/mail"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
	"trekkstay/utils"
)

func NewHotelEmpHandler(db *database.Database, requestValidator *validator.Validate) hotelEmpHandler.HotelEmpHandler {
	// Config
	mailConfig := config.LoadConfig(&models.MailConfig{}).(*models.MailConfig)
	jwtConfig := config.LoadConfig(&models.JWTConfig{}).(*models.JWTConfig)

	// HotelEmp Repository
	hotelEmpRepoReader := repository.NewHotelEmpReaderRepository(*db)
	hotelEmpRepoWriter := repository.NewHotelEmpWriterRepository(*db)

	hashAlgo := utils.NewHashAlgo()
	jwtToken := jwt.NewJWT(jwtConfig.JWTSecretKey)
	mailer := mail.NewMailer(mailConfig)

	return hotelEmpHandler.NewHotelEmpHandler(
		requestValidator,
		usecase.NewCreateHotelEmpUseCase(mailer, hashAlgo, hotelEmpRepoReader, hotelEmpRepoWriter),
		usecase.NewCreateHotelOwnerUseCase(hashAlgo, hotelEmpRepoReader, hotelEmpRepoWriter),
		usecase.NewLoginHotelEmpUseCase(jwtToken, jwtConfig.AccessTokenExpiry,
			jwtConfig.RefreshTokenExpiry, hashAlgo, hotelEmpRepoReader),
	)
}

func (r *RouteHandler) hotelEmpRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/hotel-emp",
		Routes: []route.Route{
			{
				Path:    "/create-owner",
				Method:  method.POST,
				Handler: r.HotelEmpHandler.HandleCreateHotelOwner,
			},
			{
				Path:    "/create-emp",
				Method:  method.POST,
				Handler: r.HotelEmpHandler.HandleCreateHotelEmp,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/login",
				Method:  method.POST,
				Handler: r.HotelEmpHandler.HandleLoginHotelEmp,
			},
		},
	}
}
