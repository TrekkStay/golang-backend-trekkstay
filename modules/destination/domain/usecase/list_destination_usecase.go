package usecase

import (
	"context"
	"trekkstay/modules/destination/constant"
	"trekkstay/modules/destination/domain/entity"
)

type ListDestinationUseCase interface {
	ExecuteListDestination(ctx context.Context) ([]entity.DestinationEntity, error)
}

type listDestinationUseCaseImpl struct {
	readerRepo DestinationReaderRepository
}

var _ ListDestinationUseCase = &listDestinationUseCaseImpl{}

func NewListDestinationUseCase(readerRepo DestinationReaderRepository) ListDestinationUseCase {
	return &listDestinationUseCaseImpl{
		readerRepo: readerRepo,
	}
}

func (useCase listDestinationUseCaseImpl) ExecuteListDestination(ctx context.Context) ([]entity.DestinationEntity, error) {
	destinations, err := useCase.readerRepo.FindDestinations(ctx)
	if err != nil {
		return nil, constant.ErrCanNotListDestination(err)
	}

	return destinations, nil
}
