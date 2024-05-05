package usecase

import (
	"context"
	"trekkstay/core"
	"trekkstay/modules/rating/domain/entity"
)

type CreateRatingUseCase interface {
	ExecuteCreateRating(ctx context.Context, rating entity.RatingEntity) error
}

type createRatingUseCaseImpl struct {
	ratingWriterRepo RatingWriterRepository
}

var _ CreateRatingUseCase = (*createRatingUseCaseImpl)(nil)

func (useCase createRatingUseCaseImpl) ExecuteCreateRating(ctx context.Context, rating entity.RatingEntity) error {
	requester := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester)
	rating.UserID = requester.GetUserID()

	return useCase.ratingWriterRepo.InsertRating(ctx, rating)
}
