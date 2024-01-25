package repository

import (
	"context"
	"trekkstay/modules/user/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type userWriterRepositoryImpl struct {
	db database.Database
}

var _ UserWriterRepository = (*userWriterRepositoryImpl)(nil)

func NewUserWriterRepository(db database.Database) UserWriterRepository {
	return &userWriterRepositoryImpl{
		db: db,
	}
}

func (repo userWriterRepositoryImpl) InsertUser(_ context.Context, userEntity entity.UserEntity) error {
	return repo.db.Executor.
		Create(&userEntity).Error
}

func (repo userWriterRepositoryImpl) DeleteUser(_ context.Context, userID string) error {
	return repo.db.Executor.
		Where("id = ?", userID).
		Delete(&entity.UserEntity{}).Error
}

func (repo userWriterRepositoryImpl) UpdateUser(_ context.Context, userEntity entity.UserEntity) error {
	return repo.db.Executor.
		Model(&entity.UserEntity{}).
		Where("id = ? OR email = ?", userEntity.ID, userEntity.Email).
		Updates(&userEntity).Error
}
