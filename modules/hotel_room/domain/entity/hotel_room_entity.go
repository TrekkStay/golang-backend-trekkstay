package entity

import "trekkstay/core"

type HotelRoomEntity struct {
	core.BaseEntity
	HotelID       string `json:"hotel_id" gorm:"column:hotel_id;not null"`
	Type          string `json:"type" gorm:"column:type;not null"`
	Description   string `json:"description" gorm:"column:description;default:null"`
	Quantity      int    `json:"quantity" gorm:"column:quantity;not null;default:1"`
	DiscountRate  int    `json:"discount_rate" gorm:"column:discount_rate;not null;default:0"`
	OriginalPrice int    `json:"original_price" gorm:"column:original_price;not null"`
	// JSONB
	Videos     MediaJSON               `json:"videos" gorm:"type:jsonb;default:null"`
	Images     MediaJSON               `json:"images" gorm:"type:jsonb;default:null"`
	Facilities HotelRoomFacilitiesJSON `json:"facilities" gorm:"type:jsonb;default:null"`
}

func (HotelRoomEntity) TableName() string {
	return "hotel_rooms"
}
