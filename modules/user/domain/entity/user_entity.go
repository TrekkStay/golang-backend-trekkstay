package entity

import "trekkstay/core"

type UserEntity struct {
	core.Entity `json:",inline"`
	FullName    string `json:"full_name" gorm:"not null;"`
	Email       string `json:"email" gorm:"uniqueIndex;not null;"`
	Phone       string `json:"phone" gorm:"uniqueIndex;default:null"`
	Status      string `json:"status" gorm:"not null;default:unverified"`
	OTP         string `json:"-" gorm:"default:null"`
	Password    string `json:"-" gorm:"not null;"`
}

func (UserEntity) TableName() string {
	return "users"
}
