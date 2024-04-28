package usecase

import (
	"context"
	"trekkstay/core"
	"trekkstay/modules/reservation/domain/entity"
)

type CreateReservationUseCase interface {
	ExecuteCreateReservation(ctx context.Context, reservation *entity.ReservationEntity) (*entity.ReservationEntity, error)
}

type createReservationUseCaseImpl struct {
	roomReaderRepo        HotelRoomReaderRepository
	reservationWriterRepo ReservationWriterRepository
	//qrCodeGenerator       QRCodeGenerator
}

var _ CreateReservationUseCase = (*createReservationUseCaseImpl)(nil)

func NewCreateReservationUseCase(
	roomReaderRepo HotelRoomReaderRepository,
	reservationWriterRepo ReservationWriterRepository,
) CreateReservationUseCase {
	return &createReservationUseCaseImpl{
		roomReaderRepo:        roomReaderRepo,
		reservationWriterRepo: reservationWriterRepo,
	}
}

func (useCase createReservationUseCaseImpl) ExecuteCreateReservation(ctx context.Context,
	reservation *entity.ReservationEntity) (*entity.ReservationEntity, error) {
	// Find room
	room, err := useCase.roomReaderRepo.FindHotelRoomByCondition(ctx,
		map[string]interface{}{"id": reservation.RoomID})

	if err != nil {
		return nil, err
	}

	// Retrieve room information
	(*reservation).Room.HotelID = (*room).HotelID
	(*reservation).Room.Type = (*room).Type
	(*reservation).Room.OriginalPrice = (*room).OriginalPrice
	(*reservation).Room.Images = entity.MediaJSON((*room).Images)
	(*reservation).Room.BookingPrice = (*room).OriginalPrice - ((*room).OriginalPrice * (*room).DiscountRate / 100)

	(*reservation).Status = "UPCOMING"
	(*reservation).UserID = ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserID()
	(*reservation).TotalPrice = (*reservation).Room.BookingPrice * (*reservation).Quantity

	if err := useCase.reservationWriterRepo.InsertReservation(ctx, reservation); err != nil {
		return nil, err
	}

	return reservation, nil
}
