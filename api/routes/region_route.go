package routes

import (
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

func (r *RouteHandler) regionRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/region",
		Routes: []route.Route{
			{
				Path:    "/list-province",
				Method:  method.GET,
				Handler: r.RegionHandler.HandleListProvince,
			},
			{
				Path:    "/list-district",
				Method:  method.GET,
				Handler: r.RegionHandler.HandleListDistrict,
			},
			{
				Path:    "/list-ward",
				Method:  method.GET,
				Handler: r.RegionHandler.HandleListWard,
			},
		},
	}
}
