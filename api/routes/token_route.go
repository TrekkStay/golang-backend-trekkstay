package routes

import (
	"trekkstay/api/middlewares"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

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
