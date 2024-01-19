package routes

import (
	"trekkstay/modules/user/api/handler"
	"trekkstay/pkgs/transport/http/route"
)

type RouteHandler struct {
	UserHandler handler.UserHandler
}

func (r *RouteHandler) InitRoutes() []route.Route {
	return []route.Route{}
}

func (r *RouteHandler) InitGroupRoutes() []route.GroupRoute {
	var routeGroup []route.GroupRoute
	routeGroup = append(routeGroup, r.userRoute())
	return routeGroup
}
