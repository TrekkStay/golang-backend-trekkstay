package repository

import (
	"context"
	"trekkstay/modules/hotel_emp/domain/entity"
	database "trekkstay/pkgs/dbs/postgres"
)

type hotelEmpReaderRepositoryImpl struct {
	db database.Database
}

var _ HotelEmpReaderRepository = (*hotelEmpReaderRepositoryImpl)(nil)

func NewHotelEmpRepoReader(db database.Database) HotelEmpReaderRepository {
	return &hotelEmpReaderRepositoryImpl{
		db: db,
	}
}

func (repo hotelEmpReaderRepositoryImpl) FindHotelEmpByCondition(ctx context.Context,
	condition map[string]interface{}) (*entity.HotelEmpEntity, error) {
	var hotelEmp entity.HotelEmpEntity

	if err := repo.db.Executor.
		WithContext(ctx).
		Where(condition).
		First(&hotelEmp).Error; err != nil {
		return nil, err
	}

	return &hotelEmp, nil
}

func (repo hotelEmpReaderRepositoryImpl) FindHotelEmpByHotelID(ctx context.Context,
	hotelID string) ([]entity.HotelEmpEntity, error) {
	var hotelEmployees []entity.HotelEmpEntity

	if err := repo.db.Executor.
		WithContext(ctx).
		Where("hotel_id = ?", hotelID).
		Find(&hotelEmployees).Error; err != nil {
		return nil, err
	}

	return hotelEmployees, nil
}
