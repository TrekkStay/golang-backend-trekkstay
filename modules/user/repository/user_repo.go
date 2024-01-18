package repository

import (
	"context"
	"trekkstay/modules/user/domain/entity"
)

type userWriRepository interface {
	FindUserByCondition(ctx context.Context, condition map[string]interface{}) (*entity.UserEntity, error)
}

type userWriterRepository interface {
	InsertUser(ctx context.Context, userEntity entity.UserEntity) error
	DeleteUser(ctx context.Context, userId string) error
	UpdateUser(ctx context.Context, userEntity entity.UserEntity) error
}

type UserRepository struct {
	userWriRepository
	userWriterRepository
}

func NewUserRepository(
	userReaderRepository userWriRepository,
	userWriterRepository userWriterRepository,
) *UserRepository {
	return &UserRepository{
		userWriRepository:    userReaderRepository,
		userWriterRepository: userWriterRepository,
	}
}
