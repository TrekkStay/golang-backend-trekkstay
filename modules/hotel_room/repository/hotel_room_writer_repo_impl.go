package repository

import (
	"context"
	"trekkstay/modules/hotel_room/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type hotelRoomWriterRepositoryImpl struct {
	db database.Database
}

var _ HotelRoomWriterRepository = (*hotelRoomWriterRepositoryImpl)(nil)

func NewHotelRoomWriterRepository(db database.Database) HotelRoomWriterRepository {
	return &hotelRoomWriterRepositoryImpl{
		db: db,
	}
}

func (repo hotelRoomWriterRepositoryImpl) InsertHotelRoom(ctx context.Context, room entity.HotelRoomEntity) error {
	return repo.db.Executor.
		WithContext(ctx).
		Create(&room).Error
}

func (repo hotelRoomWriterRepositoryImpl) UpdateHotelRoom(ctx context.Context, room entity.HotelRoomEntity) error {
	return repo.db.Executor.
		WithContext(ctx).
		Model(&room).
		Where("id = ?", room.ID).
		Updates(room).Error
}

func (repo hotelRoomWriterRepositoryImpl) DeleteHotelRoom(ctx context.Context, roomID string) error {
	return repo.db.Executor.
		WithContext(ctx).
		Where("id = ?", roomID).
		Delete(&entity.HotelRoomEntity{}).Error
}
