package routes

import (
	"trekkstay/api/middlewares"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

func (r *RouteHandler) hotelRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/hotel",
		Routes: []route.Route{
			{
				Path:    "/create",
				Method:  method.POST,
				Handler: r.HotelHandler.HandleCreatHotel,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
		},
	}
}
