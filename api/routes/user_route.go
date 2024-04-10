package routes

import (
	"github.com/go-playground/validator/v10"
	"trekkstay/api/middlewares"
	"trekkstay/config"
	"trekkstay/config/models"
	userHandler "trekkstay/modules/user/api/handler"
	"trekkstay/modules/user/domain/usecase"
	"trekkstay/modules/user/repository"
	database "trekkstay/pkgs/dbs/postgres"
	"trekkstay/pkgs/jwt"
	"trekkstay/pkgs/mail"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
	"trekkstay/utils"
)

func NewUserHandler(db *database.Database, requestValidator *validator.Validate) userHandler.UserHandler {
	// Config
	mailConfig := config.LoadConfig(&models.MailConfig{}).(*models.MailConfig)
	jwtConfig := config.LoadConfig(&models.JWTConfig{}).(*models.JWTConfig)

	// User Repository
	userRepoReader := repository.NewUserReaderRepository(*db)
	userRepoWriter := repository.NewUserWriterRepository(*db)

	hashAlgo := utils.NewHashAlgo()
	jwtToken := jwt.NewJWT(jwtConfig.JWTSecretKey)
	mailer := mail.NewMailer(mailConfig)

	return userHandler.NewUserHandler(requestValidator,
		usecase.NewCreateUserUseCase(hashAlgo, userRepoReader, userRepoWriter),
		usecase.NewUpdateUserUseCase(userRepoReader, userRepoWriter),
		usecase.NewLoginUserUseCase(jwtToken, jwtConfig.AccessTokenExpiry,
			jwtConfig.RefreshTokenExpiry, hashAlgo, userRepoReader),
		usecase.NewChangePasswordUseCase(hashAlgo, userRepoReader, userRepoWriter),
		usecase.NewResetPasswordUseCase(mailer, hashAlgo, userRepoReader, userRepoWriter),
	)
}

func (r *RouteHandler) userRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/user",
		Routes: []route.Route{
			{
				Path:    "/signup",
				Method:  method.POST,
				Handler: r.UserHandler.HandleCreateUser,
			},
			{
				Path:    "/update",
				Method:  method.PATCH,
				Handler: r.UserHandler.HandleUpdateUser,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/login",
				Method:  method.POST,
				Handler: r.UserHandler.HandleLoginUser,
			},
			{
				Path:    "/change-password",
				Method:  method.POST,
				Handler: r.UserHandler.HandleChangePassword,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/reset-password",
				Method:  method.POST,
				Handler: r.UserHandler.HandleResetPassword,
			},
		},
	}
}
