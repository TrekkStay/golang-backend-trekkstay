package usecase

import (
	"context"
	"trekkstay/core"
	"trekkstay/modules/hotel/domain/entity"
	emp "trekkstay/modules/hotel_emp/domain/entity"
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

type hotelEmpReaderRepository interface {
	FindHotelEmpByCondition(ctx context.Context, condition map[string]interface{}) (*emp.HotelEmpEntity, error)
}
