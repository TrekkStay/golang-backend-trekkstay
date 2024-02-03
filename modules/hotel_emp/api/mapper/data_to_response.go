package mapper

import (
	"trekkstay/modules/hotel_emp/api/model/res"
	"trekkstay/modules/hotel_emp/domain/entity"
)

func CovertUserEntityToLoginHotelEmpRes(entity entity.HotelEmpEntity) res.LoginHotelEmpRes {
	return res.LoginHotelEmpRes{
		HotelID:  entity.HotelID,
		FullName: entity.FullName,
		Email:    entity.Email,
		Phone:    entity.Phone,
		Role:     entity.Role,
		Token: res.Token{
			AccessToken:  entity.AccessToken,
			RefreshToken: entity.RefreshToken,
		},
	}
}
