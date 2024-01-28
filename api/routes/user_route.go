package routes

import (
	"trekkstay/api/middlewares"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

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
				Path:    "/forgot-password",
				Method:  method.POST,
				Handler: r.UserHandler.HandleForgotPassword,
			},
			{
				Path:    "/refresh-token",
				Method:  method.GET,
				Handler: r.UserHandler.HandleRefreshToken,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
		},
	}
}
