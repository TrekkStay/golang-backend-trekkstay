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
