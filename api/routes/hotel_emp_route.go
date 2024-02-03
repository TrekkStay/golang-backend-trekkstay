package routes

import (
	"trekkstay/api/middlewares"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

func (r *RouteHandler) hotelEmpRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/hotel-emp",
		Routes: []route.Route{
			{
				Path:    "/create-owner",
				Method:  method.POST,
				Handler: r.HotelEmpHandler.HandleCreateHotelOwner,
			},
			{
				Path:    "/create-emp",
				Method:  method.POST,
				Handler: r.HotelEmpHandler.HandleCreateHotelEmp,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/login",
				Method:  method.POST,
				Handler: r.HotelEmpHandler.HandleLoginHotelEmp,
			},
		},
	}
}
