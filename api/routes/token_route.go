package routes

import (
	"trekkstay/api/middlewares"
	"trekkstay/config"
	"trekkstay/config/models"
	tokenHandler "trekkstay/modules/token/api/handler"
	"trekkstay/modules/token/domain/usecase"
	"trekkstay/pkgs/jwt"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

func NewTokenHandler() tokenHandler.TokenHandler {
	jwtConfig := config.LoadConfig(&models.JWTConfig{}).(*models.JWTConfig)
	jwtToken := jwt.NewJWT(jwtConfig.JWTSecretKey)

	return tokenHandler.NewTokenHandler(
		usecase.NewRefreshTokenUseCase(jwtToken, jwtConfig.AccessTokenExpiry, jwtConfig.RefreshTokenExpiry),
	)
}

func (r *RouteHandler) tokenRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/token",
		Routes: []route.Route{
			{
				Path:    "/refresh-token",
				Method:  method.GET,
				Handler: r.TokenHandler.HandleRefreshToken,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
		},
	}
}
