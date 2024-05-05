package usecase

import (
	"context"
	"trekkstay/modules/hotel_room/domain/entity"
)

type hotelRoomReaderRepository interface {
	FindHotelRooms(ctx context.Context, filter entity.HotelRoomFilterEntity) ([]entity.HotelRoomEntity, error)
	FindHotelRoomByCondition(ctx context.Context, condition map[string]interface{}) (*entity.HotelRoomEntity, error)
}

type hotelRoomWriterRepository interface {
	InsertHotelRoom(ctx context.Context, room *entity.HotelRoomEntity) error
	UpdateHotelRoom(ctx context.Context, room entity.HotelRoomEntity) error
	DeleteHotelRoom(ctx context.Context, roomID string) error
}
