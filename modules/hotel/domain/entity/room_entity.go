package entity

import "trekkstay/core"

type RoomEntity struct {
	core.Entity
	HotelID     string      `json:"hotel_id" gorm:"column:hotel_id"`
	Hotel       HotelEntity `json:"hotel" gorm:"foreignKey:HotelID;references:ID"`
	Name        string      `json:"name" gorm:"column:name"`
	View        string      `json:"view" gorm:"column:view;default:none"` // none, city_view, sea_view, mountain_view
	Price       int         `json:"price" gorm:"column:price"`
	RoomSize    int         `json:"room_size" gorm:"column:room_size"`
	NumOfAdult  int         `json:"number_of_adult" gorm:"column:number_of_adult;default:1"`
	NumberOfBed int         `json:"number_of_bed" gorm:"column:number_of_bed;default:1"`
	Balcony     bool        `json:"balcony" gorm:"column:balcony;default:false"`
	Shower      bool        `json:"shower" gorm:"column:shower;default:true"`
	BathTub     bool        `json:"bath_tub" gorm:"column:bath_tub;default:false"`
}

func (RoomEntity) TableName() string {
	return "rooms"
}
