package usecase

import (
	"context"
	"errors"
	"log/slog"
	"trekkstay/core"
	"trekkstay/modules/hotel/constant"
	"trekkstay/modules/hotel/domain/entity"
	"trekkstay/pkgs/log"
)

type CreateHotelUseCase interface {
	ExecuteCreateHotel(ctx context.Context, hotelEntity *entity.HotelEntity) error
}

type createHotelUseCaseImpl struct {
	readerRepo hotelReaderRepository
	writerRepo hotelWriterRepository
}

var _ CreateHotelUseCase = (*createHotelUseCaseImpl)(nil)

func NewCreateHotelUseCase(readerRepo hotelReaderRepository, writerRepo hotelWriterRepository) CreateHotelUseCase {
	return &createHotelUseCaseImpl{
		readerRepo: readerRepo,
		writerRepo: writerRepo,
	}
}

func (useCase createHotelUseCaseImpl) ExecuteCreateHotel(ctx context.Context, hotelEntity *entity.HotelEntity) error {
	requester := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester)

	// Check permission
	if requester.GetRole() != "OWNER" {
		err := errors.New("permission denied")
		log.JsonLogger.Error("ExecuteCreateHotel.check_permission",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrCanNotCreateHotel(err)
	}

	// Find hotel by condition
	hotelEntity.OwnerID = requester.GetUserID()
	hotel, _ := useCase.readerRepo.FindHotelByCondition(ctx, map[string]interface{}{
		"owner_id": hotelEntity.OwnerID,
	})
	if hotel != nil {
		err := errors.New("one account only have one hotel")
		log.JsonLogger.Error("ExecuteCreateHotel.find_hotel_by_condition",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrCanNotCreateHotel(err)
	}

	if err := useCase.writerRepo.InsertHotel(ctx, hotelEntity); err != nil {
		log.JsonLogger.Error("ExecuteCreateHotel.insert_hotel",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrCanNotCreateHotel(err)
	}

	return nil
}
