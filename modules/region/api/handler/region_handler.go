package handler

import (
	"github.com/gin-gonic/gin"
	"trekkstay/modules/region/domain/usecase"
)

type RegionHandler interface {
	HandleListProvince(c *gin.Context)
	HandleListDistrict(c *gin.Context)
	HandleListWard(c *gin.Context)
}

type regionHandler struct {
	listProvinceUseCase usecase.ListProvinceUseCase
	listDistrictUseCase usecase.ListDistrictUseCase
	listWardUseCase     usecase.ListWardUseCase
}

func NewRegionHandler(
	listProvinceUseCase usecase.ListProvinceUseCase,
	listDistrictUseCase usecase.ListDistrictUseCase,
	listWardUseCase usecase.ListWardUseCase,
) RegionHandler {
	return &regionHandler{
		listProvinceUseCase: listProvinceUseCase,
		listDistrictUseCase: listDistrictUseCase,
		listWardUseCase:     listWardUseCase,
	}
}
