package usecase

import (
	"context"
	"trekkstay/modules/destination/constant"
	"trekkstay/modules/destination/domain/entity"
)

type CreateDestinationUseCase interface {
	ExecuteCreateDestination(ctx context.Context, destination entity.DestinationEntity) error
}

type createDestinationUseCaseImpl struct {
	writerRepo DestinationWriterRepository
}

var _ CreateDestinationUseCase = &createDestinationUseCaseImpl{}

func NewCreateDestinationUseCase(writerRepo DestinationWriterRepository) CreateDestinationUseCase {
	return &createDestinationUseCaseImpl{
		writerRepo: writerRepo,
	}
}

func (useCase createDestinationUseCaseImpl) ExecuteCreateDestination(ctx context.Context, destination entity.DestinationEntity) error {
	err := useCase.writerRepo.InsertDestination(ctx, destination)
	if err != nil {
		return constant.ErrCanNotCreateDestination(err)
	}

	return nil
}
