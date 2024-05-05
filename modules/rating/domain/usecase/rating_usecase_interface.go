package usecase

import (
	"context"
	"trekkstay/modules/rating/domain/entity"
)

type RatingReaderRepository interface {
	CountRatingAndAveragePoint(ctx context.Context, hotelID string) (int64, int64, error)
	FindRating(ctx context.Context, filter entity.RatingEntity) ([]entity.RatingEntity, error)
}

type RatingWriterRepository interface {
	InsertRating(ctx context.Context, rating entity.RatingEntity) error
}
