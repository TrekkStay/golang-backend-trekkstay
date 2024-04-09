package routes

import (
	"trekkstay/api/middlewares"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

func (r *RouteHandler) hotelRoomRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/hotel-room",
		Routes: []route.Route{
			{
				Path:    "/create",
				Method:  method.POST,
				Handler: r.HotelRoomHandler.HandleCreateHotelRoom,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/filter",
				Method:  method.GET,
				Handler: r.HotelRoomHandler.HandleFilterHotelRoom,
			},
		},
	}
}
