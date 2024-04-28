package usecase

import (
	"context"
	"trekkstay/modules/payment/domain/entity"
)

type paymentReaderRepository interface {
	FindPaymentByReservationID(ctx context.Context, reservationID string) (*entity.PaymentEntity, error)
}

type paymentWriterRepository interface {
	InsertPayment(ctx context.Context, payment *entity.PaymentEntity) error
	UpdatePaymentStatus(ctx context.Context, paymentID string, status string) error
}
