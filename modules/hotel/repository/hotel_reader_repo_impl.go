package repository

import (
	"context"
	"gorm.io/gorm"
	"trekkstay/core"
	"trekkstay/modules/hotel/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type hotelReaderRepositoryImpl struct {
	db database.Database
}

var _ HotelReaderRepository = (*hotelReaderRepositoryImpl)(nil)

func NewHotelRepoReader(db database.Database) HotelReaderRepository {
	return &hotelReaderRepositoryImpl{
		db: db,
	}
}

func (repo hotelReaderRepositoryImpl) FindHotelByID(ctx context.Context, hotelID string) (*entity.HotelEntity, error) {
	var hotelEntity entity.HotelEntity

	if err := repo.db.Executor.
		WithContext(ctx).
		Where("id = ?", hotelID).
		First(&hotelEntity).Error; err != nil {
		return nil, err
	}

	return &hotelEntity, nil
}

func (repo hotelReaderRepositoryImpl) FindHotels(ctx context.Context,
	filter entity.HotelFilterEntity, page, limit int) (*core.Pagination, error) {
	var paging core.Pagination
	var hotels []entity.HotelEntity

	paging.Limit = limit
	paging.Page = page

	var scopeFunctions []func(d *gorm.DB) *gorm.DB

	if filter.WardCode != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("ward_code = ?", *filter.WardCode)
		})
	}

	if filter.DistrictCode != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("district_code = ?", *filter.DistrictCode)
		})
	}

	if filter.ProvinceCode != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("province_code = ?", *filter.ProvinceCode)
		})
	}

	if filter.PriceOrder != nil {
		if *filter.PriceOrder == "asc" {
			paging.Sort = "min_price ASC, hotels.created_at DESC"
		} else {
			paging.Sort = "min_price DESC, hotels.created_at DESC"
		}
	}

	tx := repo.db.Executor.WithContext(ctx).Scopes(scopeFunctions...)
	txTotalRows := tx.Model(&entity.HotelEntity{}).Scopes(scopeFunctions...)
	result := tx.
		Select("hotels.*, MIN(rooms.original_price) as min_price").
		Scopes(core.Paginate(&paging, txTotalRows)).
		Preload("Rooms").
		Preload("Province").
		Preload("District").
		Preload("Ward").
		Preload("HotelFacility").
		InnerJoins("join rooms on rooms.hotel_id = hotels.id").
		Group("hotels.id").
		Find(&hotels)

	paging.Rows = hotels

	return &paging, result.Error
}
