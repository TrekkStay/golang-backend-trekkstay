package repository

import (
	"context"
	"trekkstay/core"
	"trekkstay/modules/reservation/domain/entity"
)

type ReservationReaderRepository interface {
	FindReservationByID(ctx context.Context, reservationID string) (*entity.ReservationEntity, error)
	FilterReservation(ctx context.Context, filter entity.ReservationFilterEntity, page, limit int) (*core.Pagination, error)
}

type ReservationWriterRepository interface {
	InsertReservation(ctx context.Context, reservation *entity.ReservationEntity) error
	UpdateReservationStatus(ctx context.Context, reservationID string, status string) error
	UpdateReservation(ctx context.Context, reservation entity.ReservationEntity) error
}
