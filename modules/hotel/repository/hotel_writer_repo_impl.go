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

func NewHotelRepoWriter(db database.Database) HotelWriterRepository {
	return &hotelWriterRepositoryImpl{db: db}
}

// InsertHotel implements HotelWriterRepository interface
func (repo hotelWriterRepositoryImpl) InsertHotel(ctx context.Context, hotel entity.HotelEntity) error {
	// Begin a transaction
	tx := repo.db.Executor.Begin().WithContext(ctx)

	// Defer a function to roll back the transaction if an error occurs
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Insert hotel
	if err := tx.Omit("Rooms", "Owner", "HotelFacility").Create(&hotel).Error; err != nil {
		// Roll back the transaction in case of an error
		tx.Rollback()
		return err
	}

	// Insert facilities
	facilities := hotel.HotelFacility
	facilities.HotelID = hotel.ID

	if err := tx.Create(&facilities).Error; err != nil {
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

// InsertRoom implements HotelWriterRepository interface
func (repo hotelWriterRepositoryImpl) InsertRoom(ctx context.Context, room entity.RoomEntity) error {
	// Begin a transaction
	tx := repo.db.Executor.Begin().WithContext(ctx)

	// Defer a function to roll back the transaction if an error occurs
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Insert room
	if err := tx.Omit("RoomFacility").Create(&room).Error; err != nil {
		// Roll back the transaction in case of an error
		tx.Rollback()
		return err
	}

	// Insert facilities
	facilities := room.RoomFacility
	facilities.RoomID = room.ID

	if err := tx.Create(&facilities).Error; err != nil {
		// Roll back the transaction in case of an error
		tx.Rollback()
		return err
	}

	// Commit the transaction if everything is successful
	tx.Commit()

	// If any panic occurred, it will be caught by the deferred function, and the transaction will be rolled back
	return nil
}

// InsertHotelEmployee implements HotelWriterRepository interface
func (repo hotelWriterRepositoryImpl) InsertHotelEmployee(ctx context.Context,
	hotelEmp entity.HotelEmployeeEntity) error {
	return repo.db.Executor.
		WithContext(ctx).
		Create(&hotelEmp).Error
}

// UpdateHotel implements HotelWriterRepository interface
func (repo hotelWriterRepositoryImpl) UpdateHotel(ctx context.Context, hotel entity.HotelEntity) error {
	// Begin a transactionB
	tx := repo.db.Executor.Begin().WithContext(ctx)

	// Defer a function to roll back the transaction if an error occurs
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Update hotel
	if err := tx.Omit("Rooms", "Owner", "HotelFacility").Updates(&hotel).Error; err != nil {
		// Roll back the transaction in case of an error
		tx.Rollback()
		return err
	}

	// Update facilities
	facilities := hotel.HotelFacility
	facilities.HotelID = hotel.ID

	if err := tx.Where("hotel_id = ?", hotel.ID).Updates(&facilities).Error; err != nil {
		// Roll back the transaction in case of an error
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

// UpdateRoom implements HotelWriterRepository interface
func (repo hotelWriterRepositoryImpl) UpdateRoom(ctx context.Context, room entity.RoomEntity) error {
	// Begin a transaction
	tx := repo.db.Executor.Begin().WithContext(ctx)

	// Defer a function to roll back the transaction if an error occurs
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Update room
	if err := tx.Omit("RoomFacility").Updates(&room).Error; err != nil {
		// Roll back the transaction in case of an error
		tx.Rollback()
		return err
	}

	// Update facilities
	facilities := room.RoomFacility
	facilities.RoomID = room.ID

	if err := tx.Where("room_id = ?", room.ID).Updates(&facilities).Error; err != nil {
		// Roll back the transaction in case of an error
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

// UpdateHotelEmployee implements HotelWriterRepository interface
func (repo hotelWriterRepositoryImpl) UpdateHotelEmployee(ctx context.Context,
	hotelEmp entity.HotelEmployeeEntity) error {
	return repo.db.Executor.
		WithContext(ctx).
		Model(&hotelEmp).
		Updates(&hotelEmp).Error
}

// DeleteHotel implements HotelWriterRepository interface
func (repo hotelWriterRepositoryImpl) DeleteHotel(ctx context.Context, hotelID string) error {
	return repo.db.Executor.
		WithContext(ctx).
		Where("id = ?", hotelID).
		Delete(&entity.HotelEntity{}).Error
}

// DeleteRoom implements HotelWriterRepository interface
func (repo hotelWriterRepositoryImpl) DeleteRoom(ctx context.Context, roomID string) error {
	return repo.db.Executor.
		WithContext(ctx).
		Where("id = ?", roomID).
		Delete(&entity.RoomEntity{}).Error
}

// DeleteHotelEmployee implements HotelWriterRepository interface
func (repo hotelWriterRepositoryImpl) DeleteHotelEmployee(ctx context.Context, employeeID string) error {
	return repo.db.Executor.
		WithContext(ctx).
		Where("id = ?", employeeID).
		Delete(&entity.HotelEmployeeEntity{}).Error
}
