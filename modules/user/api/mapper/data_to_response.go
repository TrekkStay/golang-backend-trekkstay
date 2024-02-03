package mapper

import (
	"trekkstay/modules/user/api/model/res"
	"trekkstay/modules/user/domain/entity"
)

func CovertUserEntityToLoginUserRes(entity entity.UserEntity) res.LoginUserRes {
	return res.LoginUserRes{
		FullName: entity.FullName,
		Email:    entity.Email,
		Phone:    entity.Phone,
		Token: res.Token{
			AccessToken:  entity.AccessToken,
			RefreshToken: entity.RefreshToken,
		},
	}
}

func ConvertTokenToResponse(accessToken string, refreshToken string) res.TokenRes {
	return res.TokenRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
