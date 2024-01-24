package routes

import (
	regionHandler "trekkstay/modules/region/api/handler"
	userHandler "trekkstay/modules/user/api/handler"
	"trekkstay/pkgs/transport/http/route"
)

type RouteHandler struct {
	UserHandler   userHandler.UserHandler
	RegionHandler regionHandler.RegionHandler
}

func (r *RouteHandler) InitRoutes() []route.Route {
	return []route.Route{}
}

func (r *RouteHandler) InitGroupRoutes() []route.GroupRoute {
	var routeGroup []route.GroupRoute
	routeGroup = append(routeGroup, r.userRoute())
	routeGroup = append(routeGroup, r.regionRoute())

	return routeGroup
}
