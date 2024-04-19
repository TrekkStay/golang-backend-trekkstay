package mapper

import (
	"trekkstay/modules/attraction/api/model/req"
	"trekkstay/modules/attraction/domain/entity"
)

func ConvertCreateAttractionReqToEntity(req req.CreateAttractionReq) entity.AttractionEntity {
	return entity.AttractionEntity{
		Name:         req.Name,
		ProvinceCode: req.ProvinceCode,
		DistrictCode: req.DistrictCode,
		WardCode:     req.WardCode,
	}
}

func ConvertFilterAttractionReqToEntity(req req.FilterAttractionReq) entity.FilterAttractionEntity {
	return entity.FilterAttractionEntity{
		ProvinceCode: req.ProvinceCode,
		DistrictCode: req.DistrictCode,
		WardCode:     req.WardCode,
	}
}
