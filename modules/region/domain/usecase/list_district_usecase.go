package usecase

import (
	"context"
	"log/slog"
	"trekkstay/modules/region/constant"
	"trekkstay/modules/region/domain/entity"
	"trekkstay/pkgs/log"
)

type ListDistrictUseCase interface {
	ExecuteListDistrict(ctx context.Context, provinceCode string) ([]entity.DistrictEntity, error)
}

type listDistrictUseCaseImpl struct {
	readerRepo RegionReaderRepository
}

var _ ListDistrictUseCase = (*listDistrictUseCaseImpl)(nil)

func NewListDistrictUseCase(readerRepo RegionReaderRepository) ListDistrictUseCase {
	return &listDistrictUseCaseImpl{
		readerRepo: readerRepo,
	}
}

func (l listDistrictUseCaseImpl) ExecuteListDistrict(ctx context.Context, provinceCode string) ([]entity.DistrictEntity, error) {
	districts, err := l.readerRepo.ListDistricts(ctx, provinceCode)
	if err != nil {
		log.JsonLogger.Error("ExecuteListDistrict.list_districts_failed",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return nil, constant.ErrCannotGetDistricts(err)
	}

	return districts, nil
}
