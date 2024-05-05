package usecase

import (
	"context"
	"trekkstay/modules/rating/domain/entity"
)

type RatingReaderRepository interface {
	FindRating(ctx context.Context, filter entity.RatingEntity) ([]entity.RatingEntity, error)
}

type RatingWriterRepository interface {
	InsertRating(ctx context.Context, rating entity.RatingEntity) error
}
