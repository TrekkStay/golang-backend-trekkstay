package usecase

import (
	"context"
	"trekkstay/modules/hotel/constant"
	"trekkstay/modules/hotel/domain/entity"
	"trekkstay/utils"
)

type GetNearMeHotelUseCase interface {
	ExecGetNearMeHotelUseCase(ctx context.Context, lat float64, lng float64, maxDistance float64) ([]entity.HotelEntity, error)
}

type getNearMeHotelUseCaseImpl struct {
	readerRepo hotelReaderRepository
}

var _ GetNearMeHotelUseCase = (*getNearMeHotelUseCaseImpl)(nil)

func NewGetNearMeHotelUseCase(readerRepo hotelReaderRepository) GetNearMeHotelUseCase {
	return &getNearMeHotelUseCaseImpl{
		readerRepo: readerRepo,
	}
}

func (useCase getNearMeHotelUseCaseImpl) ExecGetNearMeHotelUseCase(ctx context.Context,
	lat float64, lng float64, maxDistance float64) ([]entity.HotelEntity, error) {
	hotels, err := useCase.readerRepo.FindHotels(ctx, entity.HotelFilterEntity{}, 1, 10000)

	if err != nil {
		return nil, constant.ErrCantNotGetHotel(err)
	}

	var nearMeHotels []entity.HotelEntity

	for _, hotel := range hotels.Rows.([]entity.HotelEntity) {
		if utils.HaversineDistance(lat, lng, hotel.Coordinates.Lat, hotel.Coordinates.Lng) <= maxDistance {
			nearMeHotels = append(nearMeHotels, hotel)
		}
	}

	return nearMeHotels, nil
}
