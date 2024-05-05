package usecase

import (
	"context"
	"trekkstay/modules/payment/domain/entity"
)

type paymentWriterRepository interface {
	InsertPayment(ctx context.Context, payment *entity.PaymentEntity) error
	UpdatePaymentStatus(ctx context.Context, paymentID string, status string) error
}
