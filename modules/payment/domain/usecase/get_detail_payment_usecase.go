package usecase

import (
	"context"
	"trekkstay/modules/payment/domain/entity"
)

type GetDetailPaymentUseCase interface {
	ExecuteGetDetailPayment(ctx context.Context, reservationID string) (*entity.PaymentEntity, error)
}

type getDetailPaymentUseCaseImpl struct {
	paymentReaderRepository paymentReaderRepository
}

var _ GetDetailPaymentUseCase = &getDetailPaymentUseCaseImpl{}

func NewGetDetailPaymentUseCase(paymentReaderRepository paymentReaderRepository) GetDetailPaymentUseCase {
	return &getDetailPaymentUseCaseImpl{paymentReaderRepository: paymentReaderRepository}
}

func (useCase getDetailPaymentUseCaseImpl) ExecuteGetDetailPayment(ctx context.Context, reservationID string) (*entity.PaymentEntity, error) {
	payment, err := useCase.paymentReaderRepository.FindPaymentByReservationID(ctx, reservationID)

	if err != nil {
		return nil, err
	}

	return payment, nil
}
