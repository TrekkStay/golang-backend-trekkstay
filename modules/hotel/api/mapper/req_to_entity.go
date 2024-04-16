package mapper

import (
	"trekkstay/core"
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
			URLS: req.Videos.URLS,
		},
		Images: entity.MediaJSON{
			URLS: req.Images.URLS,
		},
	}
}

func ConvertUpdateHotelReqToEntity(req req.UpdateHotelReq) entity.HotelEntity {
	return entity.HotelEntity{
		BaseEntity: core.BaseEntity{
			ID: req.ID,
		},
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
			URLS: req.Videos.URLS,
		},
		Images: entity.MediaJSON{
			URLS: req.Images.URLS,
		},
	}
}

func ConvertFilterHotelReqToEntity(req req.FilterHotelReq) entity.HotelFilterEntity {
	return entity.HotelFilterEntity{
		Name:         req.Name,
		ProvinceCode: req.ProvinceCode,
		DistrictCode: req.DistrictCode,
		WardCode:     req.WardCode,
		PriceOrder:   req.PriceOrder,
	}
}
