package repository

import (
	"context"
	"trekkstay/modules/user/domain/entity"
	database "trekkstay/pkgs/db"
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

func (repo userWriterRepositoryImpl) InsertUser(ctx context.Context, userEntity entity.UserEntity) error {
	return repo.db.Executor.
		Create(&userEntity).Error
}

func (repo userWriterRepositoryImpl) DeleteUser(ctx context.Context, userId string) error {
	return repo.db.Executor.
		Where("id = ?", userId).
		Delete(&entity.UserEntity{}).Error
}

func (repo userWriterRepositoryImpl) UpdateUser(ctx context.Context, userEntity entity.UserEntity) error {
	return repo.db.Executor.
		Model(&entity.UserEntity{}).
		Where("id = ? OR email = ?", userEntity.Id, userEntity.Email).
		Updates(&userEntity).Error
}
