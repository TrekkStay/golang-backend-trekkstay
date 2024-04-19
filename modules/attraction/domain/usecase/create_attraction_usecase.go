package usecase

import (
	"context"
	"trekkstay/modules/attraction/constant"
	"trekkstay/modules/attraction/domain/entity"
)

type CreateAttractionUseCase interface {
	ExecuteCreateAttraction(ctx context.Context, attraction entity.AttractionEntity) error
}

type createAttractionUseCaseImpl struct {
	writerRepo attractionWriterRepository
}

var _ CreateAttractionUseCase = &createAttractionUseCaseImpl{}

func NewCreateAttractionUseCase(writerRepo attractionWriterRepository) CreateAttractionUseCase {
	return &createAttractionUseCaseImpl{
		writerRepo: writerRepo,
	}
}

func (useCase createAttractionUseCaseImpl) ExecuteCreateAttraction(ctx context.Context, attraction entity.AttractionEntity) error {
	err := useCase.writerRepo.InsertAttraction(ctx, attraction)
	if err != nil {
		return constant.ErrCanNotCreateAttraction(err)
	}

	return nil
}
