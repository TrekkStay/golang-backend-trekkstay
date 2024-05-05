package repository

import (
	"context"
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

	for i := range rating {
		var user entity.UserEntity

		repo.db.Executor.WithContext(ctx).Model(&entity.UserEntity{}).Where("id = ?", rating[i].UserID).First(&user)
		rating[i].User = user
	}

	return rating, nil
}
