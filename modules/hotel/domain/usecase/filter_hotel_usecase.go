package usecase

import (
	"context"
	"log/slog"
	"trekkstay/core"
	"trekkstay/modules/hotel/constant"
	"trekkstay/modules/hotel/domain/entity"
	"trekkstay/pkgs/log"
)

type FilterHotelUseCase interface {
	FilterHotel(ctx context.Context, filter entity.HotelFilterEntity, page, limit int) (*core.Pagination, error)
}

type filterHotelUseCaseImpl struct {
	readerRepo hotelReaderRepository
}

var _ FilterHotelUseCase = (*filterHotelUseCaseImpl)(nil)

func NewFilterHotelUseCase(readerRepo hotelReaderRepository) FilterHotelUseCase {
	return &filterHotelUseCaseImpl{
		readerRepo: readerRepo,
	}
}

func (useCase filterHotelUseCaseImpl) FilterHotel(ctx context.Context,
	filter entity.HotelFilterEntity, page, limit int) (*core.Pagination, error) {
	hotels, err := useCase.readerRepo.FindHotels(ctx, filter, page, limit)
	if err != nil {
		log.JsonLogger.Error("FilterHotel.find_hotels",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return nil, constant.ErrCantNotGetHotel(err)
	}

	return hotels, nil
}
