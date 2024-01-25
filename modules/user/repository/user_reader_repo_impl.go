package repository

import (
	"context"
	"trekkstay/modules/user/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type userReaderRepositoryImpl struct {
	db database.Database
}

var _ UserReaderRepository = (*userReaderRepositoryImpl)(nil)

func NewUserReaderRepository(db database.Database) UserReaderRepository {
	return &userReaderRepositoryImpl{
		db: db,
	}
}

func (repo userReaderRepositoryImpl) FindUserByCondition(_ context.Context,
	condition map[string]interface{}) (*entity.UserEntity, error) {
	var userEntity entity.UserEntity

	err := repo.db.Executor.Where(condition).First(&userEntity).Error
	if err != nil {
		return nil, err
	}

	return &userEntity, nil
}
