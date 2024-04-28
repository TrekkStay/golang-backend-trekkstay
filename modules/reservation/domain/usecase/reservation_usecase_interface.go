package usecase

import (
	"context"
	"trekkstay/core"
	room "trekkstay/modules/hotel_room/domain/entity"
	"trekkstay/modules/reservation/domain/entity"
)

type ReservationReaderRepository interface {
	FindReservationByID(ctx context.Context, reservationID string) (*entity.ReservationEntity, error)
	FilterReservation(ctx context.Context, filter entity.ReservationFilterEntity, page, limit int) (*core.Pagination, error)
}

type ReservationWriterRepository interface {
	InsertReservation(ctx context.Context, reservation *entity.ReservationEntity) error
	UpdateReservationStatus(ctx context.Context, reservationID string, status string) error
}

type HotelRoomReaderRepository interface {
	FindHotelRooms(ctx context.Context, filter room.HotelRoomFilterEntity) ([]room.HotelRoomEntity, error)
	FindHotelRoomByCondition(ctx context.Context, condition map[string]interface{}) (*room.HotelRoomEntity, error)
}
