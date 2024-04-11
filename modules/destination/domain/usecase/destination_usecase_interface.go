package usecase

import (
	"context"
	"trekkstay/modules/destination/domain/entity"
)

type DestinationReaderRepository interface {
	FindDestinations(ctx context.Context) ([]entity.DestinationEntity, error)
}

type DestinationWriterRepository interface {
	InsertDestination(ctx context.Context, destination entity.DestinationEntity) error
}
