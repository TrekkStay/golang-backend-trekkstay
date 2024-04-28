package repository

import (
	"context"
	"trekkstay/modules/reservation/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type reservationWriterRepositoryImpl struct {
	db database.Database
}

var _ ReservationWriterRepository = (*reservationWriterRepositoryImpl)(nil)

func NewReservationWriterRepository(db database.Database) ReservationWriterRepository {
	return &reservationWriterRepositoryImpl{db: db}
}

func (repo reservationWriterRepositoryImpl) InsertReservation(ctx context.Context, reservation *entity.ReservationEntity) error {
	return repo.db.Executor.
		WithContext(ctx).
		Create(&reservation).Error
}

func (repo reservationWriterRepositoryImpl) UpdateReservationStatus(ctx context.Context, reservationID string, status string) error {
	return repo.db.Executor.
		WithContext(ctx).
		Model(&entity.ReservationEntity{}).
		Where("id = ?", reservationID).
		Update("status", status).Error
}
