package usecase

import (
	"context"
	"trekkstay/modules/hotel_emp/domain/entity"
)

type hotelEmpReaderRepository interface {
	FindHotelEmpByCondition(ctx context.Context, condition map[string]interface{}) (*entity.HotelEmpEntity, error)
	FindHotelEmpByHotelID(ctx context.Context, hotelID string) ([]entity.HotelEmpEntity, error)
}

type hotelEmpWriterRepository interface {
	InsertHotelEmp(ctx context.Context, hotelEmp entity.HotelEmpEntity) error
	UpdateHotelEmp(ctx context.Context, hotelEmp entity.HotelEmpEntity) error
	DeleteHotelEmp(ctx context.Context, employeeID string) error
}

type TokenProvider interface {
	Generate(payload map[string]interface{}, expiry int) (map[string]interface{}, error)
}

type HashAlgo interface {
	HashAndSalt(pwd []byte) (string, error)
	ComparePasswords(hashedPwd string, plainPwd []byte) error
}

type Mailer interface {
	SendMail(to, subject, templatePath string, data interface{}) error
}
