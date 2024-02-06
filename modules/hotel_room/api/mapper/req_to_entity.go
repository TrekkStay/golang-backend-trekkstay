package mapper

import (
	"trekkstay/modules/hotel_room/api/model/req"
	"trekkstay/modules/hotel_room/domain/entity"
)

func ConvertCreateHotelRoomReqToEntity(req req.CreateHotelRoomReq) entity.HotelRoomEntity {
	return entity.HotelRoomEntity{
		HotelID:       req.HotelID,
		Type:          req.Type,
		Description:   req.Description,
		Quantity:      req.Quantity,
		DiscountRate:  req.DiscountRate,
		OriginalPrice: req.OriginalPrice,
		Videos:        entity.MediaJSON{URLS: req.Videos.URLS},
		Images:        entity.MediaJSON{URLS: req.Images.URLS},
		Facilities: entity.HotelRoomFacilitiesJSON{
			RoomSize:       req.Facilities.RoomSize,
			NumberOfBed:    req.Facilities.NumberOfBed,
			View:           req.Facilities.View,
			Balcony:        req.Facilities.Balcony,
			BathTub:        req.Facilities.BathTub,
			Kitchen:        req.Facilities.Kitchen,
			Television:     req.Facilities.Television,
			Shower:         req.Facilities.Shower,
			NonSmoking:     req.Facilities.NonSmoking,
			HairDryer:      req.Facilities.HairDryer,
			AirConditioner: req.Facilities.AirConditioner,
			Slippers:       req.Facilities.Slippers,
			Sleeps: entity.SleepJSON{
				Adults:   req.Facilities.Sleeps.Adults,
				Children: req.Facilities.Sleeps.Children,
			},
		},
	}
}
