package repository

import (
	"context"
	"gorm.io/gorm"
	"trekkstay/modules/attraction/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type attractionReaderRepositoryImpl struct {
	db database.Database
}

var _ AttractionReaderRepository = (*attractionReaderRepositoryImpl)(nil)

func NewAttractionReaderRepository(db database.Database) AttractionReaderRepository {
	return &attractionReaderRepositoryImpl{
		db: db,
	}
}

func (repo attractionReaderRepositoryImpl) FindAttractions(ctx context.Context, filter entity.FilterAttractionEntity) ([]entity.AttractionEntity, error) {
	var destinations []entity.AttractionEntity

	var scopeFunctions []func(d *gorm.DB) *gorm.DB

	if filter.ProvinceCode != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("province_code = ?", *filter.ProvinceCode)
		})
	}

	if filter.DistrictCode != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("district_code = ?", *filter.DistrictCode)
		})
	}

	if filter.WardCode != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("ward_code = ?", *filter.WardCode)
		})
	}

	tx := repo.db.Executor.WithContext(ctx).Scopes(scopeFunctions...)
	txTotalRows := tx.Model(&entity.AttractionEntity{}).Scopes(scopeFunctions...)

	if err := txTotalRows.
		Find(&destinations).Error; err != nil {
		return nil, err
	}

	return destinations, nil
}
