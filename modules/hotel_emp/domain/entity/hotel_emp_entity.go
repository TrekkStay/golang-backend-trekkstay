package entity

import "trekkstay/core"

type HotelEmpEntity struct {
	core.BaseEntity `json:",inline"`
	HotelID         string `json:"hotel_id" gorm:"default:null"`
	FullName        string `json:"full_name" gorm:"not null;"`
	Email           string `json:"email" gorm:"uniqueIndex;not null;"`
	Phone           string `json:"phone" gorm:"uniqueIndex;default:null"`
	Contract        string `json:"contract" gorm:"default:full_time"`
	BaseSalary      int    `json:"base_salary" gorm:"default:null"`
	Role            string `json:"role" gorm:"not null;default:employee"`
	Status          string `json:"status" gorm:"not null;default:unverified"`
	OTP             string `json:"-" gorm:"default:null"`
	Password        string `json:"-" gorm:"not null;"`
	AccessToken     string `json:"-" gorm:"-"`
	RefreshToken    string `json:"-" gorm:"-"`
}

func (HotelEmpEntity) TableName() string {
	return "hotel_employees"
}
