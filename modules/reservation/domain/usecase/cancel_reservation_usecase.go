package usecase

import "context"

type CancelReservationUseCase interface {
	ExecuteCancelReservation(ctx context.Context, reservationID string) error
}

type cancelReservationUseCaseImpl struct {
	writerRepo ReservationWriterRepository
}

var _ CancelReservationUseCase = (*cancelReservationUseCaseImpl)(nil)

func NewCancelReservationUseCase(writerRepo ReservationWriterRepository) CancelReservationUseCase {
	return &cancelReservationUseCaseImpl{writerRepo: writerRepo}
}

func (useCase cancelReservationUseCaseImpl) ExecuteCancelReservation(ctx context.Context, reservationID string) error {
	return useCase.writerRepo.UpdateReservationStatus(ctx, reservationID, "CANCELLED")
}
