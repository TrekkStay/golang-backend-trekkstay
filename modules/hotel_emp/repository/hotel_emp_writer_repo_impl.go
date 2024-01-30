package repository

import (
	"context"
	"trekkstay/modules/hotel_emp/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type hotelEmpWriterRepositoryImpl struct {
	db database.Database
}

var _ HotelEmpWriterRepository = (*hotelEmpWriterRepositoryImpl)(nil)

func NewHotelEmpRepoWriter(db database.Database) HotelEmpWriterRepository {
	return &hotelEmpWriterRepositoryImpl{
		db: db,
	}
}

func (repo hotelEmpWriterRepositoryImpl) InsertHotelEmp(ctx context.Context, hotelEmp entity.HotelEmpEntity) error {
	return repo.db.Executor.
		WithContext(ctx).
		Create(&hotelEmp).Error
}

func (repo hotelEmpWriterRepositoryImpl) UpdateHotelEmp(ctx context.Context, hotelEmp entity.HotelEmpEntity) error {
	return repo.db.Executor.
		WithContext(ctx).
		Model(&hotelEmp).
		Updates(&hotelEmp).Error
}

func (repo hotelEmpWriterRepositoryImpl) DeleteHotelEmp(ctx context.Context, employeeID string) error {
	return repo.db.Executor.
		WithContext(ctx).
		Where("id = ?", employeeID).
		Delete(&entity.HotelEmpEntity{}).Error
}
