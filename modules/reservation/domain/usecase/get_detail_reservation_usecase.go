package usecase

import (
	"context"
	"trekkstay/modules/reservation/domain/entity"
)

type GetDetailReservationUseCase interface {
	ExecuteGetDetailReservation(ctx context.Context, reservationID string) (*entity.ReservationEntity, error)
}

type getDetailReservationUseCaseImpl struct {
	reservationReaderRepository ReservationReaderRepository
}

var _ GetDetailReservationUseCase = &getDetailReservationUseCaseImpl{}

func NewGetDetailReservationUseCase(reservationReaderRepository ReservationReaderRepository) GetDetailReservationUseCase {
	return &getDetailReservationUseCaseImpl{
		reservationReaderRepository: reservationReaderRepository,
	}
}

func (useCase getDetailReservationUseCaseImpl) ExecuteGetDetailReservation(ctx context.Context,
	reservationID string) (*entity.ReservationEntity, error) {
	reservation, err := useCase.reservationReaderRepository.FindReservationByID(ctx, reservationID)
	if err != nil {
		return nil, err
	}

	return reservation, nil
}
