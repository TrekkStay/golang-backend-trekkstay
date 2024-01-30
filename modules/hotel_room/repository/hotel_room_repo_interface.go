package repository

import (
	"context"
	"trekkstay/modules/hotel_room/domain/entity"
)

type HotelRoomReaderRepository interface {
	FindHotelRooms(ctx context.Context, filter entity.HotelRoomFilterEntity) ([]entity.HotelRoomEntity, error)
}

type HotelRoomWriterRepository interface {
	InsertHotelRoom(ctx context.Context, room entity.HotelRoomEntity) error
	UpdateHotelRoom(ctx context.Context, room entity.HotelRoomEntity) error
	DeleteHotelRoom(ctx context.Context, roomID string) error
}
