package repository

import (
	"context"
	"trekkstay/modules/payment/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type paymentWriterRepositoryImpl struct {
	db database.Database
}

var _ PaymentWriterRepository = (*paymentWriterRepositoryImpl)(nil)

func NewPaymentWriterRepository(db database.Database) PaymentWriterRepository {
	return &paymentWriterRepositoryImpl{db: db}
}

func (repo paymentWriterRepositoryImpl) InsertPayment(ctx context.Context, payment *entity.PaymentEntity) error {
	if err := repo.db.Executor.
		WithContext(ctx).
		Create(&payment).Error; err != nil {
		return err
	}

	return nil
}

func (repo paymentWriterRepositoryImpl) UpdatePaymentStatus(ctx context.Context, paymentID string, status string) error {
	if err := repo.db.Executor.
		WithContext(ctx).
		Model(&entity.PaymentEntity{}).
		Where("id = ?", paymentID).
		Update("status", status).Error; err != nil {
		return err
	}

	return nil
}
