package routes

import (
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

func (r *RouteHandler) uploadRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/upload",
		Routes: []route.Route{
			{
				Path:    "/media",
				Method:  method.POST,
				Handler: r.UploadHandler.HandleUploadMedia,
			},
		},
	}
}
