package usecase

import (
	"context"
	"trekkstay/modules/attraction/domain/entity"
)

type attractionReaderRepository interface {
	FindAttractions(ctx context.Context, filter entity.FilterAttractionEntity) ([]entity.AttractionEntity, error)
}

type attractionWriterRepository interface {
	InsertAttraction(ctx context.Context, attraction entity.AttractionEntity) error
}
