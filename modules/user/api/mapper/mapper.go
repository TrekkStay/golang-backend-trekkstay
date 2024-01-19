package mapper

import (
	"trekkstay/modules/user/api/model/req"
	"trekkstay/modules/user/domain/entity"
)

func ConvertCreateUserReqToUserEntity(req req.CreateUserReq) entity.UserEntity {
	return entity.UserEntity{
		FullName: req.FullName,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
		Status:   entity.UNVERIFIED.Value(),
	}
}
