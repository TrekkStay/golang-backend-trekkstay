package usecase

import (
	"context"
	"log/slog"
	"trekkstay/modules/hotel_room/constant"
	"trekkstay/modules/hotel_room/domain/entity"
	"trekkstay/pkgs/log"
)

type FilterHotelRoomUseCase interface {
	ExecuteFilterHotelRoom(ctx context.Context, filter entity.HotelRoomFilterEntity) ([]entity.HotelRoomEntity, error)
}

type filterHotelRoomUseCaseImpl struct {
	readerRepo hotelRoomReaderRepository
}

var _ FilterHotelRoomUseCase = (*filterHotelRoomUseCaseImpl)(nil)

func NewFilterHotelRoomUseCase(readerRepo hotelRoomReaderRepository) FilterHotelRoomUseCase {
	return &filterHotelRoomUseCaseImpl{
		readerRepo: readerRepo,
	}
}

func (useCase filterHotelRoomUseCaseImpl) ExecuteFilterHotelRoom(ctx context.Context,
	filter entity.HotelRoomFilterEntity) ([]entity.HotelRoomEntity, error) {
	if filter.HotelID == nil {
		return nil, constant.ErrHotelIDIsRequired(nil)
	}

	rooms, err := useCase.readerRepo.FindHotelRooms(ctx, filter)

	if err != nil {
		log.JsonLogger.Error("ExecuteFilterHotelRoom.find_hotel_rooms",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return nil, constant.ErrCanNotFilterHotelRoom(err)
	}

	return rooms, nil
}
