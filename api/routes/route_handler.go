package routes

import (
	"github.com/go-playground/validator/v10"
	"trekkstay/config"
	"trekkstay/config/models"
	regionHandler "trekkstay/modules/region/api/handler"
	regionUseCase "trekkstay/modules/region/domain/usecase"
	regionRepo "trekkstay/modules/region/repository"
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
	UserHandler   userHandler.UserHandler
	RegionHandler regionHandler.RegionHandler
}

func (r *RouteHandler) InitGroupRoutes() []route.GroupRoute {
	var routeGroup []route.GroupRoute
	routeGroup = append(routeGroup, r.userRoute())
	routeGroup = append(routeGroup, r.regionRoute())

	return routeGroup
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
		userUseCase.NewForgotPasswordUseCase(mailer, hashAlgo, userRepoReader, userRepoWriter),
		userUseCase.NewRefreshTokenUseCase(jwtToken, jwtConfig.AccessTokenExpiry, jwtConfig.RefreshTokenExpiry),
	)
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
