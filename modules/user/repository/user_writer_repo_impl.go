package repository

import (
	"context"
	"trekkstay/modules/user/domain/entity"
	database "trekkstay/pkgs/db"
)

type UserWriterRepositoryImpl struct {
	db database.Database
}

var _ userWriterRepository = (*UserWriterRepositoryImpl)(nil)

func NewUserWriterRepository(db database.Database) *UserWriterRepositoryImpl {
	return &UserWriterRepositoryImpl{
		db: db,
	}
}

func (repo UserWriterRepositoryImpl) InsertUser(ctx context.Context, userEntity entity.UserEntity) error {
	return repo.db.Executor.
		Create(&userEntity).Error
}

func (repo UserWriterRepositoryImpl) DeleteUser(ctx context.Context, userId string) error {
	return repo.db.Executor.
		Where("id = ?", userId).
		Delete(&entity.UserEntity{}).Error
}

func (repo UserWriterRepositoryImpl) UpdateUser(ctx context.Context, userEntity entity.UserEntity) error {
	return repo.db.Executor.
		Model(&entity.UserEntity{}).
		Where("id = ? OR email = ?", userEntity.Id, userEntity.Email).
		Updates(&userEntity).Error
}
