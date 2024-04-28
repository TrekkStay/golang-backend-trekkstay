package mapper

import (
	"trekkstay/modules/attraction/api/model/req"
	"trekkstay/modules/attraction/domain/entity"
)

func ConvertCreateAttractionReqToEntity(req req.CreateAttractionReq) entity.AttractionEntity {
	return entity.AttractionEntity{
		Name:         req.Name,
		Lat:          req.Lat,
		Lng:          req.Lng,
		ProvinceCode: req.ProvinceCode,
		DistrictCode: req.DistrictCode,
		WardCode:     req.WardCode,
	}
}

func ConvertFilterAttractionReqToEntity(req req.FilterAttractionReq) entity.FilterAttractionEntity {
	if len(req.LocationCode) == 2 {
		return entity.FilterAttractionEntity{
			ProvinceCode: &req.LocationCode,
		}
	}

	if len(req.LocationCode) == 3 {
		return entity.FilterAttractionEntity{
			DistrictCode: &req.LocationCode,
		}
	}

	return entity.FilterAttractionEntity{
		WardCode: &req.LocationCode,
	}
}
