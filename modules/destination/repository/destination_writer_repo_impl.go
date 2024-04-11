package repository

import (
	"context"
	"trekkstay/modules/destination/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type destinationWriterRepositoryImpl struct {
	db database.Database
}

var _ DestinationWriterRepository = (*destinationWriterRepositoryImpl)(nil)

func NewDestinationWriterRepository(db database.Database) DestinationWriterRepository {
	return &destinationWriterRepositoryImpl{
		db: db,
	}
}

func (repo destinationWriterRepositoryImpl) InsertDestination(ctx context.Context, destination entity.DestinationEntity) error {
	if err := repo.db.Executor.
		WithContext(ctx).
		Create(&destination).Error; err != nil {
		return err
	}

	return nil
}
