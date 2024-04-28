package entity

import "trekkstay/core"

type ReservationEntity struct {
	core.BaseEntity
	RoomID       string        `json:"room_id" gorm:"column:room_id;"`
	UserID       string        `json:"user_id" gorm:"column:user_id;"`
	QRCodeURL    string        `json:"qr_code_url" gorm:"column:qr_code_url;default:null"`
	Quantity     int           `json:"quantity" gorm:"column:quantity;"`
	TotalPrice   int           `json:"total_price" gorm:"column:total_price;"`
	CheckInDate  string        `json:"check_in_date" gorm:"column:check_in_date;"`
	CheckOutDate string        `json:"check_out_date" gorm:"column:check_out_date;"`
	Status       string        `json:"status" gorm:"column:status;default:UPCOMING;"`
	GuestInfo    GuestInfoJSON `json:"guest_info" gorm:"type:json;column:guest_info;"`
	Room         RoomJSON      `json:"room" gorm:"type:json;column:room;"`
	User         UserEntity    `json:"user" gorm:"foreignKey:UserID;references:ID;"`
}

func (ReservationEntity) TableName() string {
	return "reservations"
}
