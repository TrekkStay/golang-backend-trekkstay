package usecase

import (
	"context"
	"trekkstay/core"
	"trekkstay/modules/hotel/constant"
	"trekkstay/modules/hotel/domain/entity"
)

type GetMyHotelUseCase interface {
	ExecuteGetMyHotel(ctx context.Context) (*entity.HotelEntity, error)
}

type getMyHotelUseCaseImpl struct {
	readerRepo hotelReaderRepository
}

var _ GetMyHotelUseCase = (*getMyHotelUseCaseImpl)(nil)

func NewGetMyHotelUseCase(readerRepo hotelReaderRepository) GetMyHotelUseCase {
	return &getMyHotelUseCaseImpl{
		readerRepo: readerRepo,
	}
}

func (useCase getMyHotelUseCaseImpl) ExecuteGetMyHotel(ctx context.Context) (*entity.HotelEntity, error) {
	requester := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester)

	hotel, err := useCase.readerRepo.FindHotelByCondition(ctx, map[string]interface{}{
		"owner_id": requester.GetUserID(),
	})

	if err != nil {
		return nil, constant.ErrCantNotGetHotel(err)
	}

	return hotel, nil
}
