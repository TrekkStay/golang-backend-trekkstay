package usecase

import (
	"context"
	"log/slog"
	"trekkstay/modules/hotel_room/constant"
	"trekkstay/modules/hotel_room/domain/entity"
	"trekkstay/pkgs/log"
)

type CreateHotelRoomUseCase interface {
	ExecuteCreateHotelRoom(ctx context.Context, hotelRoomEntity entity.HotelRoomEntity) error
}

type createHotelRoomUseCaseImpl struct {
	writerRepo hotelRoomWriterRepository
}

var _ CreateHotelRoomUseCase = (*createHotelRoomUseCaseImpl)(nil)

func NewCreateHotelRoomUseCase(writerRepo hotelRoomWriterRepository) CreateHotelRoomUseCase {
	return &createHotelRoomUseCaseImpl{
		writerRepo: writerRepo,
	}
}

func (useCase createHotelRoomUseCaseImpl) ExecuteCreateHotelRoom(ctx context.Context,
	hotelRoomEntity entity.HotelRoomEntity) error {
	if err := useCase.writerRepo.InsertHotelRoom(ctx, hotelRoomEntity); err != nil {
		log.JsonLogger.Error("ExecuteCreateHotelRoom.insert_hotel_room",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrCanNotCreateHotelRoom(err)
	}

	return nil
}
