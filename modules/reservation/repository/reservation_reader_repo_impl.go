package repository

import (
	"context"
	"gorm.io/gorm"
	"trekkstay/core"
	"trekkstay/modules/reservation/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type reservationReaderRepositoryImpl struct {
	db database.Database
}

var _ ReservationReaderRepository = (*reservationReaderRepositoryImpl)(nil)

func NewReservationReaderRepository(db database.Database) ReservationReaderRepository {
	return &reservationReaderRepositoryImpl{db: db}
}

func (repo reservationReaderRepositoryImpl) FindReservationByID(ctx context.Context, reservationID string) (*entity.ReservationEntity, error) {
	var reservation entity.ReservationEntity

	if err := repo.db.Executor.
		WithContext(ctx).
		Preload("User").
		Where("id = ?", reservationID).
		First(&reservation).Error; err != nil {
		return nil, err
	}

	return &reservation, nil
}

func (repo reservationReaderRepositoryImpl) FilterReservation(ctx context.Context,
	filter entity.ReservationFilterEntity, page, limit int) (*core.Pagination, error) {
	var paging core.Pagination
	var reservations []entity.ReservationEntity

	paging.Limit = limit
	paging.Page = page

	var scopeFunctions []func(d *gorm.DB) *gorm.DB

	if filter.UserID != nil && filter.HotelID == nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("user_id = ?", *filter.UserID)
		})
	}

	if filter.Status != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("status = ?", *filter.Status)
		})
	}

	if filter.HotelID != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("(room->>'hotel_id')::string = ?", *filter.HotelID)
		})
	}

	if filter.Date != nil {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("check_in_date = ?", *filter.Date)
		})
	}

	tx := repo.db.Executor.WithContext(ctx).Model(&entity.ReservationEntity{})
	txTotalRows := tx.Model(&entity.ReservationEntity{}).Scopes(scopeFunctions...)

	result := tx.
		Scopes(core.Paginate(&paging, txTotalRows)).
		Preload("User").
		Find(&reservations)

	paging.Rows = reservations

	return &paging, result.Error
}
