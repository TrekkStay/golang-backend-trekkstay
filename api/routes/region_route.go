package routes

import (
	regionHandler "trekkstay/modules/region/api/handler"
	"trekkstay/modules/region/domain/usecase"
	"trekkstay/modules/region/repository"
	database "trekkstay/pkgs/dbs/postgres"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

func NewRegionHandler(db *database.Database) regionHandler.RegionHandler {
	// Region Repository
	regionRepoReader := repository.NewRegionReaderRepository(*db)

	return regionHandler.NewRegionHandler(
		usecase.NewListProvinceUseCase(regionRepoReader),
		usecase.NewListDistrictUseCase(regionRepoReader),
		usecase.NewListWardUseCase(regionRepoReader),
	)
}

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
