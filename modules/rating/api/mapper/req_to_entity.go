package mapper

import (
	"trekkstay/modules/rating/api/model/req"
	"trekkstay/modules/rating/domain/entity"
)

func ConvertCreateRatingReqToEntity(req req.CreateRatingReq) entity.RatingEntity {
	return entity.RatingEntity{
		HotelID:        req.HotelID,
		Title:          req.Title,
		TypeOfTraveler: req.TypeOfTraveler,
		Point:          req.Point,
		Summary:        req.Summary,
	}
}
