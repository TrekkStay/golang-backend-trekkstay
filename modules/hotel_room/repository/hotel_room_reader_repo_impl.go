package repository

import (
	"context"
	"gorm.io/gorm"
	"trekkstay/modules/hotel_room/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type hotelRoomReaderRepositoryImpl struct {
	db database.Database
}

var _ HotelRoomReaderRepository = (*hotelRoomReaderRepositoryImpl)(nil)

func NewHotelRoomReaderRepository(db database.Database) HotelRoomReaderRepository {
	return &hotelRoomReaderRepositoryImpl{
		db: db,
	}
}

func (repo hotelRoomReaderRepositoryImpl) FindHotelRoomByCondition(ctx context.Context,
	condition map[string]interface{}) (*entity.HotelRoomEntity, error) {
	var room entity.HotelRoomEntity
	err := repo.db.Executor.WithContext(ctx).Where(condition).First(&room).Error
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (repo hotelRoomReaderRepositoryImpl) FindHotelRooms(ctx context.Context,
	filter entity.HotelRoomFilterEntity) ([]entity.HotelRoomEntity, error) {
	var rooms []entity.HotelRoomEntity

	var scopeFunctions []func(d *gorm.DB) *gorm.DB

	if filter.HotelID != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("hotel_id = ?", *filter.HotelID)
		})
	}

	if filter.NonSmoking != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("(facilities->>'non_smoking')::boolean = ?", *filter.NonSmoking)
		})
	}

	if filter.Balcony != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("(facilities->>'balcony')::boolean = ?", *filter.Balcony)
		})
	}

	if filter.BathTub != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("(facilities->>'bathtub')::boolean = ?", *filter.BathTub)
		})
	}

	if filter.Kitchen != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("(facilities->>'kitchen')::boolean = ?", *filter.Kitchen)
		})
	}

	if filter.PriceOrder != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Order("original_price " + *filter.PriceOrder)
		})
	}

	tx := repo.db.Executor.WithContext(ctx).Scopes(scopeFunctions...)
	txTotalRows := tx.Model(&entity.HotelRoomEntity{}).Scopes(scopeFunctions...)
	result := txTotalRows.Find(&rooms)

	return rooms, result.Error
}
