package repository

import (
	"context"
	"trekkstay/modules/user/domain/entity"
	database "trekkstay/pkgs/db"
)

type UserReaderRepositoryImpl struct {
	db database.Database
}

var _ userWriRepository = (*UserReaderRepositoryImpl)(nil)

func NewUserReaderRepository(db database.Database) *UserReaderRepositoryImpl {
	return &UserReaderRepositoryImpl{
		db: db,
	}
}

func (repo UserReaderRepositoryImpl) FindUserByCondition(ctx context.Context,
	condition map[string]interface{}) (*entity.UserEntity, error) {
	var userEntity entity.UserEntity

	err := repo.db.Executor.Where(condition).First(&userEntity).Error
	if err != nil {
		return nil, err
	}

	return &userEntity, nil
}
