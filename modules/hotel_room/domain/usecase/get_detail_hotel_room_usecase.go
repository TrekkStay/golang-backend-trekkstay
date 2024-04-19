package usecase

import (
	"context"
	"log/slog"
	"trekkstay/modules/hotel_room/domain/entity"
	"trekkstay/pkgs/log"
)

type GetDetailHotelRoomUseCase interface {
	ExecuteGetDetailHotelRoom(ctx context.Context, hotelRoomID string) (*entity.HotelRoomEntity, error)
}

type getDetailHotelRoomUseCaseImpl struct {
	readerRepo hotelRoomReaderRepository
}

var _ GetDetailHotelRoomUseCase = (*getDetailHotelRoomUseCaseImpl)(nil)

func NewGetDetailHotelRoomUseCase(readerRepo hotelRoomReaderRepository) GetDetailHotelRoomUseCase {
	return &getDetailHotelRoomUseCaseImpl{
		readerRepo: readerRepo,
	}
}

func (useCase getDetailHotelRoomUseCaseImpl) ExecuteGetDetailHotelRoom(ctx context.Context, hotelRoomID string) (*entity.HotelRoomEntity, error) {
	hotelRoom, err := useCase.readerRepo.FindHotelRoomByCondition(ctx, map[string]interface{}{
		"id": hotelRoomID,
	})
	if err != nil {
		log.JsonLogger.Error("GetDetailHotelRoom.find_hotel_rooms",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return nil, err
	}

	return hotelRoom, err
}
