package usecase

import (
	"context"
	"trekkstay/modules/hotel/domain/entity"
)

type StaticHotelUseCase interface {
	ExecuteGetStatistic(ctx context.Context, hotelID string) (*entity.HotelEntity, error)
}
