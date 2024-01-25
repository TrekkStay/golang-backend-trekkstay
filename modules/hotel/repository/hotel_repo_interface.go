package repository

import (
	"context"
	"trekkstay/core"
	"trekkstay/modules/hotel/domain/entity"
)

type HotelReaderRepository interface {
	FindHotelByID(ctx context.Context, hotelID string) (*entity.HotelEntity, error)
	FindHotels(ctx context.Context, filter entity.HotelFilterEntity, page, limit int) (*core.Pagination, error)
	FindRooms(ctx context.Context, filter entity.RoomFilterEntity) ([]entity.RoomEntity, error)
}

type HotelWriterRepository interface {
	InsertHotel(ctx context.Context, hotel entity.HotelEntity) error
	InsertRoom(ctx context.Context, room entity.RoomEntity) error
	InsertHotelEmployee(ctx context.Context, hotelEmp entity.HotelEmployeeEntity) error
	UpdateHotel(ctx context.Context, hotel entity.HotelEntity) error
	UpdateRoom(ctx context.Context, room entity.RoomEntity) error
	UpdateHotelEmployee(hotelEmp entity.HotelEmployeeEntity) error
	DeleteHotel(ctx context.Context, hotelID string) error
	DeleteRoom(ctx context.Context, roomID string) error
	DeleteHotelEmployee(ctx context.Context, employeeID string) error
}
