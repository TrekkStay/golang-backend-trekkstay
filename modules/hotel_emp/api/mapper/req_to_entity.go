package mapper

import (
	"trekkstay/modules/hotel_emp/api/model/req"
	"trekkstay/modules/hotel_emp/domain/entity"
)

func ConvertCreateHotelEmpReqEntity(req req.CreateHotelEmpReq) entity.HotelEmpEntity {
	return entity.HotelEmpEntity{
		FullName:   req.FullName,
		Email:      req.Email,
		Phone:      req.Phone,
		Contract:   req.Contract,
		BaseSalary: req.BaseSalary,
		Status:     entity.ACTIVE.Value(),
	}
}

func ConvertCreateHotelOwnerReqEntity(req req.CreateHotelOwnerReq) entity.HotelEmpEntity {
	return entity.HotelEmpEntity{
		FullName: req.FullName,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
		Status:   entity.UNVERIFIED.Value(),
	}
}
