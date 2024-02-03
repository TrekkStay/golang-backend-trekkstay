package routes

import (
	"github.com/go-playground/validator/v10"
	"trekkstay/config"
	"trekkstay/config/models"
	hotelEmpHandler "trekkstay/modules/hotel_emp/api/handler"
	hotelEmpUseCase "trekkstay/modules/hotel_emp/domain/usecase"
	hotelEmpRepo "trekkstay/modules/hotel_emp/repository"
	regionHandler "trekkstay/modules/region/api/handler"
	regionUseCase "trekkstay/modules/region/domain/usecase"
	regionRepo "trekkstay/modules/region/repository"
	tokenHandler "trekkstay/modules/token/api/handler"
	tokenUseCase "trekkstay/modules/token/domain/usecase"
	userHandler "trekkstay/modules/user/api/handler"
	userUseCase "trekkstay/modules/user/domain/usecase"
	userRepo "trekkstay/modules/user/repository"
	database "trekkstay/pkgs/dbs/postgres"
	"trekkstay/pkgs/jwt"
	"trekkstay/pkgs/mail"
	"trekkstay/pkgs/transport/http/route"
	"trekkstay/utils"
)

type RouteHandler struct {
	UserHandler     userHandler.UserHandler
	RegionHandler   regionHandler.RegionHandler
	HotelEmpHandler hotelEmpHandler.HotelEmpHandler
	TokenHandler    tokenHandler.TokenHandler
}

func (r *RouteHandler) InitGroupRoutes() []route.GroupRoute {
	var routeGroup []route.GroupRoute
	routeGroup = append(routeGroup, r.regionRoute())
	routeGroup = append(routeGroup, r.userRoute())
	routeGroup = append(routeGroup, r.hotelEmpRoute())
	routeGroup = append(routeGroup, r.tokenRoute())

	return routeGroup
}

func NewRegionHandler(db *database.Database) regionHandler.RegionHandler {
	// Region Repository
	regionRepoReader := regionRepo.NewRegionReaderRepository(*db)

	return regionHandler.NewRegionHandler(
		regionUseCase.NewListProvinceUseCase(regionRepoReader),
		regionUseCase.NewListDistrictUseCase(regionRepoReader),
		regionUseCase.NewListWardUseCase(regionRepoReader),
	)
}
func NewUserHandler(db *database.Database, requestValidator *validator.Validate) userHandler.UserHandler {
	// Config
	mailConfig := config.LoadConfig(&models.MailConfig{}).(*models.MailConfig)
	jwtConfig := config.LoadConfig(&models.JWTConfig{}).(*models.JWTConfig)

	// User Repository
	userRepoReader := userRepo.NewUserReaderRepository(*db)
	userRepoWriter := userRepo.NewUserWriterRepository(*db)

	hashAlgo := utils.NewHashAlgo()
	jwtToken := jwt.NewJWT(jwtConfig.JWTSecretKey)
	mailer := mail.NewMailer(mailConfig)

	return userHandler.NewUserHandler(requestValidator,
		userUseCase.NewCreateUserUseCase(hashAlgo, userRepoReader, userRepoWriter),
		userUseCase.NewUpdateUserUseCase(userRepoReader, userRepoWriter),
		userUseCase.NewLoginUserUseCase(jwtToken, jwtConfig.AccessTokenExpiry,
			jwtConfig.RefreshTokenExpiry, hashAlgo, userRepoReader),
		userUseCase.NewChangePasswordUseCase(hashAlgo, userRepoReader, userRepoWriter),
		userUseCase.NewResetPasswordUseCase(mailer, hashAlgo, userRepoReader, userRepoWriter),
	)
}

func NewHotelEmpHandler(db *database.Database, requestValidator *validator.Validate) hotelEmpHandler.HotelEmpHandler {
	// Config
	mailConfig := config.LoadConfig(&models.MailConfig{}).(*models.MailConfig)
	jwtConfig := config.LoadConfig(&models.JWTConfig{}).(*models.JWTConfig)

	// HotelEmp Repository
	hotelEmpRepoReader := hotelEmpRepo.NewHotelEmpReaderRepository(*db)
	hotelEmpRepoWriter := hotelEmpRepo.NewHotelEmpWriterRepository(*db)

	hashAlgo := utils.NewHashAlgo()
	jwtToken := jwt.NewJWT(jwtConfig.JWTSecretKey)
	mailer := mail.NewMailer(mailConfig)

	return hotelEmpHandler.NewHotelEmpHandler(
		requestValidator,
		hotelEmpUseCase.NewCreateHotelEmpUseCase(mailer, hashAlgo, hotelEmpRepoReader, hotelEmpRepoWriter),
		hotelEmpUseCase.NewCreateHotelOwnerUseCase(hashAlgo, hotelEmpRepoReader, hotelEmpRepoWriter),
		hotelEmpUseCase.NewLoginHotelEmpUseCase(jwtToken, jwtConfig.AccessTokenExpiry,
			jwtConfig.RefreshTokenExpiry, hashAlgo, hotelEmpRepoReader),
	)
}

func NewTokenHandler() tokenHandler.TokenHandler {
	jwtConfig := config.LoadConfig(&models.JWTConfig{}).(*models.JWTConfig)
	jwtToken := jwt.NewJWT(jwtConfig.JWTSecretKey)

	return tokenHandler.NewTokenHandler(
		tokenUseCase.NewRefreshTokenUseCase(jwtToken, jwtConfig.AccessTokenExpiry, jwtConfig.RefreshTokenExpiry),
	)
}
