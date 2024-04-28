package usecase

import (
	"context"
	"trekkstay/core"
	"trekkstay/modules/reservation/domain/entity"
)

type FilterReservationUseCase interface {
	ExecuteFilterReservation(ctx context.Context, filter entity.ReservationFilterEntity, page, limit int) (*core.Pagination, error)
}

type filterReservationUseCaseImpl struct {
	readerRepo ReservationReaderRepository
}

var _ FilterReservationUseCase = (*filterReservationUseCaseImpl)(nil)

func NewFilterReservationUseCase(readerRepo ReservationReaderRepository) FilterReservationUseCase {
	return &filterReservationUseCaseImpl{readerRepo: readerRepo}
}

func (useCase filterReservationUseCaseImpl) ExecuteFilterReservation(ctx context.Context,
	filter entity.ReservationFilterEntity, page, limit int) (*core.Pagination, error) {

	userID := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserID()

	filter.UserID = &userID

	return useCase.readerRepo.FilterReservation(ctx, filter, page, limit)
}
