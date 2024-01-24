package usecase

import (
	"context"
	"log/slog"
	"trekkstay/modules/region/constant"
	"trekkstay/modules/region/domain/entity"
	"trekkstay/pkgs/log"
)

type ListWardUseCase interface {
	ExecuteListWard(ctx context.Context, districtCode string) ([]entity.WardEntity, error)
}

type listWardUseCaseImpl struct {
	readerRepo RegionReaderRepository
}

var _ ListWardUseCase = (*listWardUseCaseImpl)(nil)

func NewListWardUseCase(readerRepo RegionReaderRepository) ListWardUseCase {
	return &listWardUseCaseImpl{
		readerRepo: readerRepo,
	}
}

func (l listWardUseCaseImpl) ExecuteListWard(ctx context.Context, districtCode string) ([]entity.WardEntity, error) {
	wards, err := l.readerRepo.ListWards(ctx, districtCode)

	if err != nil {
		log.JsonLogger.Error("ExecuteListWard.list_wards_failed",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return nil, constant.ErrCannotGetWards(err)
	}

	return wards, nil
}
