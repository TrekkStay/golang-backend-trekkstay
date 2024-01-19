package mapper

import (
	"trekkstay/modules/user/api/model/res"
	"trekkstay/modules/user/domain/entity"
)

func CovertUserEntityToLoginUserRes(entity entity.UserEntity) res.LoginUserRes {
	return res.LoginUserRes{
		FullName:     entity.FullName,
		Email:        entity.Email,
		Phone:        entity.Phone,
		Status:       entity.Status,
		AccessToken:  entity.AccessToken,
		RefreshToken: entity.RefreshToken,
	}
}
