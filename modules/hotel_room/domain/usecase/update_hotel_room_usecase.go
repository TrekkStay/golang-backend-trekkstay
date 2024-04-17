package usecase

import (
	"context"
	"trekkstay/modules/hotel_room/constant"
	"trekkstay/modules/hotel_room/domain/entity"
)

type UpdateHotelRoomUseCase interface {
	ExecUpdateHotelRoomUseCase(ctx context.Context, hotelRoom entity.HotelRoomEntity) error
}

type updateHotelRoomUseCaseImpl struct {
	writeRepo hotelRoomWriterRepository
}

var _ UpdateHotelRoomUseCase = (*updateHotelRoomUseCaseImpl)(nil)

func NewUpdateHotelRoomUseCase(writeRepo hotelRoomWriterRepository) UpdateHotelRoomUseCase {
	return &updateHotelRoomUseCaseImpl{
		writeRepo: writeRepo,
	}
}

func (useCase updateHotelRoomUseCaseImpl) ExecUpdateHotelRoomUseCase(ctx context.Context, hotelRoom entity.HotelRoomEntity) error {
	if err := useCase.writeRepo.UpdateHotelRoom(ctx, hotelRoom); err != nil {
		return constant.ErrSomethingWentWrong(nil)
	}

	return nil
}
