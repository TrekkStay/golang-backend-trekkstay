package entity

import "trekkstay/core"

type RoomEntity struct {
	core.Entity
	HotelID      string             `json:"hotel_id" gorm:"column:hotel_id;not null"`
	Type         string             `json:"type" gorm:"column:type;not null"`
	Quantity     int                `json:"quantity" gorm:"column:quantity;not null;default:1"`
	OriginPrice  int                `json:"origin_price" gorm:"column:origin_price;not null"`
	Videos       MediaObject        `json:"videos" gorm:"type:jsonb;default:null"`
	Images       MediaObject        `json:"images" gorm:"type:jsonb;default:null"`
	RoomFacility RoomFacilityEntity `json:"facility" gorm:"foreignKey:RoomID;references:ID"`
}

func (RoomEntity) TableName() string {
	return "rooms"
}
