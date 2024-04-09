package routes

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"trekkstay/config"
	"trekkstay/config/models"
	hotelHandler "trekkstay/modules/hotel/api/handler"
	hotelUseCase "trekkstay/modules/hotel/domain/usecase"
	hotelRepo "trekkstay/modules/hotel/repository"
	hotelEmpHandler "trekkstay/modules/hotel_emp/api/handler"
	hotelEmpUseCase "trekkstay/modules/hotel_emp/domain/usecase"
	hotelEmpRepo "trekkstay/modules/hotel_emp/repository"
	hotelRoomHandler "trekkstay/modules/hotel_room/api/handler"
	hotelRoomUseCase "trekkstay/modules/hotel_room/domain/usecase"
	hotelRoomRepo "trekkstay/modules/hotel_room/repository"
	regionHandler "trekkstay/modules/region/api/handler"
	regionUseCase "trekkstay/modules/region/domain/usecase"
	regionRepo "trekkstay/modules/region/repository"
	tokenHandler "trekkstay/modules/token/api/handler"
	tokenUseCase "trekkstay/modules/token/domain/usecase"
	userHandler "trekkstay/modules/user/api/handler"
	userUseCase "trekkstay/modules/user/domain/usecase"
	userRepo "trekkstay/modules/user/repository"
	database "trekkstay/pkgs/dbs/postgres"
	"trekkstay/pkgs/dbs/redis"
	"trekkstay/pkgs/jwt"
	"trekkstay/pkgs/mail"
	"trekkstay/pkgs/s3"
	"trekkstay/pkgs/transport/http/route"
	"trekkstay/utils"
)

type RouteHandler struct {
	UserHandler      userHandler.UserHandler
	RegionHandler    regionHandler.RegionHandler
	HotelEmpHandler  hotelEmpHandler.HotelEmpHandler
	HotelRoomHandler hotelRoomHandler.HotelRoomHandler
	HotelHandler     hotelHandler.HotelHandler
	TokenHandler     tokenHandler.TokenHandler
	UploadHandler    s3.UploadHandler
}

func (r *RouteHandler) InitGroupRoutes() []route.GroupRoute {
	var routeGroup []route.GroupRoute
	routeGroup = append(routeGroup, r.regionRoute())
	routeGroup = append(routeGroup, r.userRoute())
	routeGroup = append(routeGroup, r.hotelEmpRoute())
	routeGroup = append(routeGroup, r.tokenRoute())
	routeGroup = append(routeGroup, r.hotelRoute())
	routeGroup = append(routeGroup, r.uploadRoute())
	routeGroup = append(routeGroup, r.hotelRoomRoute())

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

func NewHotelRoomHandler(db *database.Database, requestValidator *validator.Validate) hotelRoomHandler.HotelRoomHandler {
	// Hotel Room Repository
	hotelRoomRepoReader := hotelRoomRepo.NewHotelRoomReaderRepository(*db)
	hotelRoomRepoWriter := hotelRoomRepo.NewHotelRoomWriterRepository(*db)

	// Redis
	redisConfig := config.LoadConfig(&models.RedisConfig{}).(*models.RedisConfig)
	var conn = redis.Connection{
		Address:  fmt.Sprint(redisConfig.RedisHost, ":", redisConfig.RedisPort),
		Password: redisConfig.RedisPassword,
		Database: redisConfig.RedisDB,
	}

	var redisInstance = redis.NewRedis(conn)

	return hotelRoomHandler.NewHotelRoomHandler(requestValidator, redisInstance,
		hotelRoomUseCase.NewCreateHotelRoomUseCase(hotelRoomRepoWriter),
		hotelRoomUseCase.NewFilterHotelRoomUseCase(hotelRoomRepoReader),
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

func NewHotelHandler(db *database.Database, requestValidator *validator.Validate) hotelHandler.HotelHandler {
	// Hotel Repository
	hotelRepoReader := hotelRepo.NewHotelReaderRepository(*db)
	hotelRepoWriter := hotelRepo.NewHotelWriterRepository(*db)

	// Redis
	redisConfig := config.LoadConfig(&models.RedisConfig{}).(*models.RedisConfig)
	var conn = redis.Connection{
		Address:  fmt.Sprint(redisConfig.RedisHost, ":", redisConfig.RedisPort),
		Password: redisConfig.RedisPassword,
		Database: redisConfig.RedisDB,
	}

	var redisInstance = redis.NewRedis(conn)

	return hotelHandler.NewHotelHandler(requestValidator, redisInstance,
		hotelUseCase.NewCreateHotelUseCase(hotelRepoReader, hotelRepoWriter),
		hotelUseCase.NewFilterHotelUseCase(hotelRepoReader),
	)
}

func NewTokenHandler() tokenHandler.TokenHandler {
	jwtConfig := config.LoadConfig(&models.JWTConfig{}).(*models.JWTConfig)
	jwtToken := jwt.NewJWT(jwtConfig.JWTSecretKey)

	return tokenHandler.NewTokenHandler(
		tokenUseCase.NewRefreshTokenUseCase(jwtToken, jwtConfig.AccessTokenExpiry, jwtConfig.RefreshTokenExpiry),
	)
}
