package repository

import (
	"context"
	"trekkstay/modules/payment/domain/entity"
)

type PaymentWriterRepository interface {
	InsertPayment(ctx context.Context, payment *entity.PaymentEntity) error
	UpdatePaymentStatus(ctx context.Context, paymentID string, status string) error
}
