package repository

import (
	"context"
	"trekkstay/modules/attraction/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type attractionWriterRepositoryImpl struct {
	db database.Database
}

var _ AttractionWriterRepository = (*attractionWriterRepositoryImpl)(nil)

func NewAttractionWriterRepository(db database.Database) AttractionWriterRepository {
	return &attractionWriterRepositoryImpl{
		db: db,
	}
}

func (repo attractionWriterRepositoryImpl) InsertAttraction(ctx context.Context, attraction entity.AttractionEntity) error {
	if err := repo.db.Executor.
		WithContext(ctx).
		Create(&attraction).Error; err != nil {
		return err
	}

	return nil
}
