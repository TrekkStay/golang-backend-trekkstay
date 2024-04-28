package repository

import (
	"context"
	"trekkstay/modules/payment/domain/entity"
)

type PaymentReaderRepository interface {
	FindPaymentByReservationID(ctx context.Context, reservationID string) (*entity.PaymentEntity, error)
}

type PaymentWriterRepository interface {
	InsertPayment(ctx context.Context, payment *entity.PaymentEntity) error
	UpdatePaymentStatus(ctx context.Context, paymentID string, status string) error
}
