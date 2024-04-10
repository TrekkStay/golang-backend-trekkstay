package usecase

import (
	"context"
	"log/slog"
	"trekkstay/modules/hotel/constant"
	"trekkstay/modules/hotel/domain/entity"
	"trekkstay/pkgs/log"
)

type GetDetailHotelUseCase interface {
	ExecuteGetDetailHotel(ctx context.Context, hotelID string) (*entity.HotelEntity, error)
}

type getDetailHotelUseCaseImpl struct {
	readerRepo hotelReaderRepository
}

var _ GetDetailHotelUseCase = (*getDetailHotelUseCaseImpl)(nil)

func NewGetDetailHotelUseCase(readerRepo hotelReaderRepository) GetDetailHotelUseCase {
	return &getDetailHotelUseCaseImpl{
		readerRepo: readerRepo,
	}
}

func (useCase getDetailHotelUseCaseImpl) ExecuteGetDetailHotel(ctx context.Context, hotelID string) (*entity.HotelEntity, error) {
	hotel, err := useCase.readerRepo.FindHotelByCondition(ctx, map[string]interface{}{
		"id": hotelID,
	})
	if err != nil {
		log.JsonLogger.Error("GetDetailHotel.find_hotels",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return nil, constant.ErrCantNotGetHotel(err)
	}

	return hotel, nil
}
