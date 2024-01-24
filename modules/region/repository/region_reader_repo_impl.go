package repository

import (
	"context"
	"trekkstay/modules/region/domain/entity"
	database "trekkstay/pkgs/db"
)

type regionReaderRepoImpl struct {
	db database.Database
}

var _ RegionReaderRepository = (*regionReaderRepoImpl)(nil)

func NewRegionReaderRepository(db database.Database) RegionReaderRepository {
	return &regionReaderRepoImpl{db: db}
}

// ListProvinces implements RegionReaderRepository.
func (repo *regionReaderRepoImpl) ListProvinces(_ context.Context) ([]entity.ProvinceEntity, error) {
	var provinces []entity.ProvinceEntity

	if err := repo.db.Executor.Find(&provinces).Error; err != nil {
		return nil, err
	}

	return provinces, nil
}

// ListDistricts implements RegionReaderRepository.
func (repo *regionReaderRepoImpl) ListDistricts(_ context.Context, provinceCode string) ([]entity.DistrictEntity, error) {
	var districts []entity.DistrictEntity

	if err := repo.db.Executor.Where("province_code = ?", provinceCode).Find(&districts).Error; err != nil {
		return nil, err
	}

	return districts, nil
}

// ListWards implements RegionReaderRepository.
func (repo *regionReaderRepoImpl) ListWards(_ context.Context, districtCode string) ([]entity.WardEntity, error) {
	var wards []entity.WardEntity

	if err := repo.db.Executor.Where("district_code = ?", districtCode).Find(&wards).Error; err != nil {
		return nil, err
	}

	return wards, nil
}
