package repository

import (
	"context"
	"trekkstay/modules/attraction/domain/entity"
)

type AttractionReaderRepository interface {
	FindAttractions(ctx context.Context, filter entity.FilterAttractionEntity) ([]entity.AttractionEntity, error)
}

type AttractionWriterRepository interface {
	InsertAttraction(ctx context.Context, attraction entity.AttractionEntity) error
}
