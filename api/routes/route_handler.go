package routes

import (
	attractionHandler "trekkstay/modules/attraction/api/handler"
	destinationHandler "trekkstay/modules/destination/api/handler"
	hotelHandler "trekkstay/modules/hotel/api/handler"
	hotelEmpHandler "trekkstay/modules/hotel_emp/api/handler"
	hotelRoomHandler "trekkstay/modules/hotel_room/api/handler"
	regionHandler "trekkstay/modules/region/api/handler"
	tokenHandler "trekkstay/modules/token/api/handler"
	userHandler "trekkstay/modules/user/api/handler"
	"trekkstay/pkgs/s3"
	"trekkstay/pkgs/transport/http/route"
)

type RouteHandler struct {
	UserHandler        userHandler.UserHandler
	RegionHandler      regionHandler.RegionHandler
	HotelEmpHandler    hotelEmpHandler.HotelEmpHandler
	HotelRoomHandler   hotelRoomHandler.HotelRoomHandler
	HotelHandler       hotelHandler.HotelHandler
	TokenHandler       tokenHandler.TokenHandler
	UploadHandler      s3.UploadHandler
	DestinationHandler destinationHandler.DestinationHandler
	AttractionHandler  attractionHandler.AttractionHandler
}

func (r *RouteHandler) InitGroupRoutes() []route.GroupRoute {
	var routeGroup []route.GroupRoute
	routeGroup = append(routeGroup, r.regionRoute())
	routeGroup = append(routeGroup, r.userRoute())
	routeGroup = append(routeGroup, r.hotelEmpRoute())
	routeGroup = append(routeGroup, r.tokenRoute())
	routeGroup = append(routeGroup, r.hotelRoute())
	routeGroup = append(routeGroup, r.uploadRoute())
	routeGroup = append(routeGroup, r.hotelRoomRoute())
	routeGroup = append(routeGroup, r.destinationRoute())
	routeGroup = append(routeGroup, r.attractionRoute())

	return routeGroup
}
