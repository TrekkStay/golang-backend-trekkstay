package usecase

import (
	"context"
	"trekkstay/modules/payment/domain/entity"
)

type CreatePaymentUseCase interface {
	ExecuteCreatePaymentUseCase(ctx context.Context, payment entity.PaymentEntity) (*entity.PaymentEntity, error)
}

type createPaymentUseCaseImpl struct {
	paymentWriterRepo paymentWriterRepository
}

var _ CreatePaymentUseCase = (*createPaymentUseCaseImpl)(nil)

func NewCreatePaymentUseCase(paymentWriterRepo paymentWriterRepository) CreatePaymentUseCase {
	return &createPaymentUseCaseImpl{paymentWriterRepo: paymentWriterRepo}
}

func (useCase createPaymentUseCaseImpl) ExecuteCreatePaymentUseCase(ctx context.Context,
	payment entity.PaymentEntity) (*entity.PaymentEntity, error) {
	if err := useCase.paymentWriterRepo.InsertPayment(ctx, &payment); err != nil {
		return nil, err
	}

	return &payment, nil
}
