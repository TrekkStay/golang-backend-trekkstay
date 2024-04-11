package repository

import (
	"context"
	"trekkstay/modules/destination/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type destinationReaderRepositoryImpl struct {
	db database.Database
}

var _ DestinationReaderRepository = (*destinationReaderRepositoryImpl)(nil)

func NewDestinationReaderRepository(db database.Database) DestinationReaderRepository {
	return &destinationReaderRepositoryImpl{
		db: db,
	}
}

func (repo destinationReaderRepositoryImpl) FindDestinations(ctx context.Context) ([]entity.DestinationEntity, error) {
	var destinations []entity.DestinationEntity

	if err := repo.db.Executor.
		WithContext(ctx).
		Find(&destinations).Error; err != nil {
		return nil, err
	}

	return destinations, nil
}
