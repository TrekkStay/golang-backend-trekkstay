package repository

import (
	"context"
	"trekkstay/modules/hotel_emp/domain/entity"
)

type HotelEmpReaderRepository interface {
	FindHotelEmpByCondition(ctx context.Context, condition map[string]interface{}) (*entity.HotelEmpEntity, error)
	FindHotelEmpByHotelID(ctx context.Context, hotelID string) ([]entity.HotelEmpEntity, error)
}

type HotelEmpWriterRepository interface {
	InsertHotelEmp(ctx context.Context, hotelEmp entity.HotelEmpEntity) error
	UpdateHotelEmp(ctx context.Context, hotelEmp entity.HotelEmpEntity) error
	DeleteHotelEmp(ctx context.Context, employeeID string) error
}
