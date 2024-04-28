package repository

import (
	"context"
	"trekkstay/modules/payment/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type paymentReaderRepositoryImpl struct {
	db database.Database
}

var _ PaymentReaderRepository = (*paymentReaderRepositoryImpl)(nil)

func NewPaymentReaderRepository(db database.Database) PaymentReaderRepository {
	return &paymentReaderRepositoryImpl{db: db}
}

func (repo paymentReaderRepositoryImpl) FindPaymentByReservationID(ctx context.Context,
	reservationID string) (*entity.PaymentEntity, error) {
	var payment entity.PaymentEntity

	if err := repo.db.Executor.
		WithContext(ctx).
		Preload("Reservation").
		Where("reservation_id = ?", reservationID).
		First(&payment).Error; err != nil {
		return nil, err
	}

	return &payment, nil
}
