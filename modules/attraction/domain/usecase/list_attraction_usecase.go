package usecase

import (
	"context"
	"trekkstay/modules/attraction/constant"
	"trekkstay/modules/attraction/domain/entity"
)

type ListAttractionUseCase interface {
	ExecuteListAttraction(ctx context.Context, filter entity.FilterAttractionEntity) ([]entity.AttractionEntity, error)
}

type listAttractionUseCaseImpl struct {
	readerRepo attractionReaderRepository
}

var _ ListAttractionUseCase = &listAttractionUseCaseImpl{}

func NewListAttractionUseCase(readerRepo attractionReaderRepository) ListAttractionUseCase {
	return &listAttractionUseCaseImpl{
		readerRepo: readerRepo,
	}
}

func (useCase listAttractionUseCaseImpl) ExecuteListAttraction(ctx context.Context,
	filter entity.FilterAttractionEntity) ([]entity.AttractionEntity, error) {
	attractions, err := useCase.readerRepo.FindAttractions(ctx, filter)
	if err != nil {
		return nil, constant.ErrCanNotListAttraction(err)
	}

	return attractions, nil
}
