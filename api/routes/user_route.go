package routes

import (
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
		},
	}
}
