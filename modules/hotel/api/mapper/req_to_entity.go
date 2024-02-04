package mapper

import (
	"trekkstay/modules/hotel/api/model/req"
	"trekkstay/modules/hotel/domain/entity"
)

func ConvertCreateHotelReqToEntity(req req.CreateHotelReq) entity.HotelEntity {
	return entity.HotelEntity{
		Name:          req.Name,
		Email:         req.Email,
		Phone:         req.Phone,
		CheckInTime:   req.CheckInTime,
		CheckOutTime:  req.CheckOutTime,
		ProvinceCode:  req.ProvinceCode,
		DistrictCode:  req.DistrictCode,
		WardCode:      req.WardCode,
		AddressDetail: req.AddressDetail,
		Description:   req.Description,
		Facilities: entity.HotelFacilitiesJSON{
			FitnessCenter:   req.Facilities.FitnessCenter,
			ConferenceRoom:  req.Facilities.ConferenceRoom,
			ParkingArea:     req.Facilities.ParkingArea,
			SwimmingPool:    req.Facilities.SwimmingPool,
			FreeWifi:        req.Facilities.FreeWifi,
			AirportTransfer: req.Facilities.AirportTransfer,
			MotorBikeRental: req.Facilities.MotorBikeRental,
			SpaService:      req.Facilities.SpaService,
			FoodService:     req.Facilities.FoodService,
			LaundryService:  req.Facilities.LaundryService,
		},
		Coordinates: entity.CoordinatesJSON{
			Lat: req.Coordinates.Lat,
			Lng: req.Coordinates.Lng,
		},
		Videos: entity.MediaJSON{
			URL: req.Videos.URL,
		},
		Images: entity.MediaJSON{
			URL: req.Images.URL,
		},
	}
}
