package usecase

import (
	"context"
	"trekkstay/modules/user/domain/entity"
)

type userReaderRepository interface {
	FindUserByCondition(ctx context.Context, condition map[string]interface{}) (*entity.UserEntity, error)
}

type userWriterRepository interface {
	InsertUser(ctx context.Context, userEntity entity.UserEntity) error
	DeleteUser(ctx context.Context, userId string) error
	UpdateUser(ctx context.Context, userEntity entity.UserEntity) error
}
