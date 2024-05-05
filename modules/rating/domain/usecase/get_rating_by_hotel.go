package usecase

import (
	"context"
	"trekkstay/modules/rating/domain/entity"
)

type GetRatingByHotelUseCase interface {
	ExecuteGetRatingByHotel(ctx context.Context, hotelID string) ([]entity.RatingEntity, error)
}

type getRatingByHotelUseCaseImpl struct {
	ratingReaderRepo RatingReaderRepository
}

var _ GetRatingByHotelUseCase = (*getRatingByHotelUseCaseImpl)(nil)

func NewGetRatingByHotelUseCase(ratingReaderRepo RatingReaderRepository) GetRatingByHotelUseCase {
	return &getRatingByHotelUseCaseImpl{ratingReaderRepo: ratingReaderRepo}
}

func (useCase getRatingByHotelUseCaseImpl) ExecuteGetRatingByHotel(ctx context.Context, hotelID string) ([]entity.RatingEntity, error) {
	return useCase.ratingReaderRepo.FindRating(ctx, entity.RatingEntity{HotelID: hotelID})
}
