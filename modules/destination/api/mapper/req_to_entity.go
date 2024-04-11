package mapper

import (
	"trekkstay/modules/destination/api/model/req"
	"trekkstay/modules/destination/domain/entity"
)

func ConvertCreateDestinationReqToEntity(req req.CreateDestinationReq) entity.DestinationEntity {
	return entity.DestinationEntity{
		Name: req.Name,
		Code: req.Code,
		Unit: req.Unit,
	}
}
