package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"trekkstay/modules/rating/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type ratingReaderRepositoryImpl struct {
	db database.Database
}

var _ RatingReaderRepository = (*ratingReaderRepositoryImpl)(nil)

func NewRatingReaderRepository(db database.Database) RatingReaderRepository {
	return &ratingReaderRepositoryImpl{db: db}
}

func (repo ratingReaderRepositoryImpl) CountRatingAndAveragePoint(ctx context.Context, hotelID string) (int64, int64, error) {
	rating, avg := int64(0), int64(0)

	if err := repo.db.Executor.
		WithContext(ctx).
		Model(&entity.RatingEntity{}).Where("hotel_id = ?", hotelID).
		Count(&rating).Select("AVG(point)").Row().Scan(&avg).Error; err != nil {
		return 0, 0, errors.New("failed to count rating and average point")
	}

	return rating, avg, nil
}

func (repo ratingReaderRepositoryImpl) FindRating(ctx context.Context, filter entity.RatingEntity) ([]entity.RatingEntity, error) {
	var scopeFunctions []func(d *gorm.DB) *gorm.DB

	if filter.HotelID != "" {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("hotel_id = ?", filter.HotelID)
		})
	}

	if filter.Point != 0 {
		scopeFunctions = append(scopeFunctions, func(d *gorm.DB) *gorm.DB {
			return d.Where("point = ?", filter.Point)
		})
	}

	tx := repo.db.Executor.WithContext(ctx).Model(&entity.RatingEntity{}).Scopes(scopeFunctions...)

	var rating []entity.RatingEntity

	if err := tx.Find(&rating).Preload("User").Error; err != nil {
		return nil, err
	}

	return rating, nil
}
