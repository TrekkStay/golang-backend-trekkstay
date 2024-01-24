package repository

import (
	"context"
	"trekkstay/modules/region/domain/entity"
)

type RegionReaderRepository interface {
	ListProvinces(ctx context.Context) ([]entity.ProvinceEntity, error)
	ListDistricts(ctx context.Context, provinceCode string) ([]entity.DistrictEntity, error)
	ListWards(ctx context.Context, districtCode string) ([]entity.WardEntity, error)
}
