package entity

import "trekkstay/core"

type RatingEntity struct {
	core.BaseEntity
	UserID         string     `json:"user_id" gorm:"not null;"`
	HotelID        string     `json:"hotel_id" gorm:"not null;"`
	Title          string     `json:"title" gorm:"not null;"`
	TypeOfTraveler string     `json:"type_of_traveler" gorm:"not null;"`
	Point          int        `json:"point" gorm:"not null;"`
	Summary        string     `json:"summary" gorm:"not null;"`
	User           UserEntity `json:"user" gorm:"foreignKey:UserID;references:ID;"`
}

func (RatingEntity) TableName() string {
	return "ratings"
}
