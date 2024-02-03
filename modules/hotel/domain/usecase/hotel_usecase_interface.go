package usecase

import (
	"context"
	"trekkstay/core"
	"trekkstay/modules/hotel/domain/entity"
)

type hotelReaderRepository interface {
	FindHotelByCondition(ctx context.Context, condition map[string]interface{}) (*entity.HotelEntity, error)
	FindHotels(ctx context.Context, filter entity.HotelFilterEntity, page, limit int) (*core.Pagination, error)
}

type hotelWriterRepository interface {
	InsertHotel(ctx context.Context, hotel entity.HotelEntity) error
	UpdateHotel(ctx context.Context, hotel entity.HotelEntity) error
	DeleteHotel(ctx context.Context, hotelID string) error
}
