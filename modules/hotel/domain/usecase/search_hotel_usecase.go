package usecase

import (
	"context"
	"trekkstay/core"
	"trekkstay/modules/hotel/domain/entity"
)

type SearchHotelUseCase interface {
	ExecuteSearchHotel(ctx context.Context, filter entity.HotelSearchEntity, page, limit int) (*core.Pagination, error)
}

type searchHotelUseCaseImpl struct {
	readerRepo hotelReaderRepository
}

var _ SearchHotelUseCase = (*searchHotelUseCaseImpl)(nil)

func NewSearchHotelUseCase(readerRepo hotelReaderRepository) SearchHotelUseCase {
	return &searchHotelUseCaseImpl{
		readerRepo: readerRepo,
	}
}

func (useCase searchHotelUseCaseImpl) ExecuteSearchHotel(ctx context.Context, filter entity.HotelSearchEntity, page, limit int) (*core.Pagination, error) {
	hotels, err := useCase.readerRepo.SearchHotel(ctx, filter, page, limit)

	if err != nil {
		return nil, err
	}

	return hotels, nil
}
