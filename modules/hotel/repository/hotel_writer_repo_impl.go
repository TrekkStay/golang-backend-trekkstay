package repository

import (
	"context"
	"trekkstay/modules/hotel/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type hotelWriterRepositoryImpl struct {
	db database.Database
}

var _ HotelWriterRepository = (*hotelWriterRepositoryImpl)(nil)

func NewHotelWriterRepository(db database.Database) HotelWriterRepository {
	return &hotelWriterRepositoryImpl{db: db}
}

// InsertHotel implements HotelWriterRepository interface
func (repo hotelWriterRepositoryImpl) InsertHotel(ctx context.Context, hotel *entity.HotelEntity) error {
	// Begin a transaction
	tx := repo.db.Executor.Begin().WithContext(ctx)

	// Defer a function to roll back the transaction if an error occurs
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Insert hotel
	if err := tx.Omit("Rooms", "Owner").Create(&hotel).Error; err != nil {
		// Roll back the transaction in case of an error
		tx.Rollback()
		return err
	}

	// Update owner
	owner := entity.HotelEmployeeEntity{}
	owner.HotelID = hotel.ID

	if err := tx.Model(&owner).Where("id = ?", hotel.OwnerID).
		Update("hotel_id", hotel.ID).Error; err != nil {
		// Roll back the transaction in case of an error
		tx.Rollback()
		return err
	}

	// Commit the transaction if everything is successful
	tx.Commit()

	// If any panic occurred, it will be caught by the deferred function, and the transaction will be rolled back
	return nil
}

// UpdateHotel implements HotelWriterRepository interface
func (repo hotelWriterRepositoryImpl) UpdateHotel(ctx context.Context, hotel entity.HotelEntity) error {
	return repo.db.Executor.
		WithContext(ctx).
		Omit("Rooms", "Owner").
		Where("id = ?", hotel.ID).
		Updates(&hotel).Error
}

// DeleteHotel implements HotelWriterRepository interface
func (repo hotelWriterRepositoryImpl) DeleteHotel(ctx context.Context, hotelID string) error {
	return repo.db.Executor.
		WithContext(ctx).
		Where("id = ?", hotelID).
		Delete(&entity.HotelEntity{}).Error
}
