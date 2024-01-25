package repository

import (
	"context"
	"gorm.io/gorm"
	"trekkstay/core"
	"trekkstay/modules/hotel/domain/entity"
	database "trekkstay/pkgs/db"
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

func (repo hotelReaderRepositoryImpl) FindHotelByID(_ context.Context, hotelID string) (*entity.HotelEntity, error) {
	var hotelEntity entity.HotelEntity

	err := repo.db.Executor.Where("id = ?", hotelID).First(&hotelEntity).Error
	if err != nil {
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
		Select("hotels.*, MIN(rooms.price) as min_price").
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

func (repo hotelReaderRepositoryImpl) FindRooms(ctx context.Context,
	filter entity.RoomFilterEntity) ([]entity.RoomEntity, error) {
	var rooms []entity.RoomEntity

	var scopeFunctions []func(d *gorm.DB) *gorm.DB

	if filter.HotelID != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("hotel_id = ?", *filter.HotelID)
		})
	}

	if filter.NonSmoking != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("room_facilities.non_smoking = ?", *filter.NonSmoking)
		})
	}

	if filter.Balcony != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("room_facilities.balcony = ?", *filter.Balcony)
		})
	}

	if filter.BathTub != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("room_facilities.bathtub = ?", *filter.BathTub)
		})
	}

	if filter.Kitchen != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("room_facilities.kitchen = ?", *filter.Kitchen)
		})
	}

	tx := repo.db.Executor.WithContext(ctx).Scopes(scopeFunctions...)
	txTotalRows := tx.Model(&entity.RoomEntity{}).Scopes(scopeFunctions...)
	result := txTotalRows.
		Joins("join room_facilities on room_facilities.room_id = rooms.id").
		Preload("RoomFacility").
		Order("origin_price ASC").
		Find(&rooms)

	return rooms, result.Error
}
