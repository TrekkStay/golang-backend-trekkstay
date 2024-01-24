package usecase

import (
	"context"
	"log/slog"
	"trekkstay/modules/region/constant"
	"trekkstay/modules/region/domain/entity"
	"trekkstay/pkgs/log"
)

type ListProvinceUseCase interface {
	ExecuteListProvince(ctx context.Context) ([]entity.ProvinceEntity, error)
}

type listProvinceUseCaseImpl struct {
	readerRepo RegionReaderRepository
}

var _ ListProvinceUseCase = (*listProvinceUseCaseImpl)(nil)

func NewListProvinceUseCase(readerRepo RegionReaderRepository) ListProvinceUseCase {
	return &listProvinceUseCaseImpl{
		readerRepo: readerRepo,
	}
}

func (l listProvinceUseCaseImpl) ExecuteListProvince(ctx context.Context) ([]entity.ProvinceEntity, error) {
	provinces, err := l.readerRepo.ListProvinces(ctx)
	if err != nil {
		log.JsonLogger.Error("ExecuteListProvince.list_provinces_failed",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return nil, constant.ErrCannotGetProvinces(err)
	}

	return provinces, nil
}
