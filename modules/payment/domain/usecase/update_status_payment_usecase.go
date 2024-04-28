package usecase

import "context"

type UpdateStatusPaymentUseCase interface {
	ExecuteUpdateStatusPayment(ctx context.Context, paymentID string, status string) error
}

type updateStatusPaymentUseCaseImpl struct {
	paymentWriterRepository paymentWriterRepository
}

var _ UpdateStatusPaymentUseCase = &updateStatusPaymentUseCaseImpl{}

func NewUpdateStatusPaymentUseCase(paymentWriterRepository paymentWriterRepository) UpdateStatusPaymentUseCase {
	return &updateStatusPaymentUseCaseImpl{paymentWriterRepository: paymentWriterRepository}
}

func (useCase updateStatusPaymentUseCaseImpl) ExecuteUpdateStatusPayment(ctx context.Context, paymentID string, status string) error {
	err := useCase.paymentWriterRepository.UpdatePaymentStatus(ctx, paymentID, status)

	if err != nil {
		return err
	}

	return nil
}
