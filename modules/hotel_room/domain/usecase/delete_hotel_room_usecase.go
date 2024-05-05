package usecase

import "context"

type DeleteHotelRoomUseCase interface {
	ExecuteDeleteHotelRoom(ctx context.Context, roomID string) error
}

type deleteHotelRoomUseCaseImpl struct {
	hotelRoomWriterRepository hotelRoomWriterRepository
}

var _ DeleteHotelRoomUseCase = &deleteHotelRoomUseCaseImpl{}

func NewDeleteHotelRoomUseCase(hotelRoomWriterRepository hotelRoomWriterRepository) DeleteHotelRoomUseCase {
	return &deleteHotelRoomUseCaseImpl{
		hotelRoomWriterRepository: hotelRoomWriterRepository,
	}
}

func (useCase deleteHotelRoomUseCaseImpl) ExecuteDeleteHotelRoom(ctx context.Context, roomID string) error {
	err := useCase.hotelRoomWriterRepository.DeleteHotelRoom(ctx, roomID)
	if err != nil {
		return err
	}

	return nil
}
