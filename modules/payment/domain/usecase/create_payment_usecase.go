package usecase

import (
	"context"
	"trekkstay/core"
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
	requester := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester)

	payment.UserID = requester.GetUserID()

	if err := useCase.paymentWriterRepo.InsertPayment(ctx, &payment); err != nil {
		return nil, err
	}

	return &payment, nil
}
