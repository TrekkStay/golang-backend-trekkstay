package repository

import (
	"context"
	"trekkstay/modules/rating/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type ratingWriterRepositoryImpl struct {
	db database.Database
}

var _ RatingWriterRepository = (*ratingWriterRepositoryImpl)(nil)

func NewRatingWriterRepository(db database.Database) RatingWriterRepository {
	return &ratingWriterRepositoryImpl{db: db}
}

func (repo ratingWriterRepositoryImpl) InsertRating(ctx context.Context, rating entity.RatingEntity) error {
	return repo.db.Executor.
		WithContext(ctx).
		Create(&rating).Error
}
