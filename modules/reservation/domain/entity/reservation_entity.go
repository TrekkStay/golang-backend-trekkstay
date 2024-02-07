package entity

import "trekkstay/core"

type ReservationEntity struct {
	core.BaseEntity
	RoomID       string        `json:"room_id" gorm:"column:room_id;"`
	UserID       string        `json:"user_id" gorm:"column:user_id;"`
	QRCodeURL    string        `json:"qr_code_url" gorm:"column:qr_code_url;"`
	PaymentID    string        `json:"payment_id" gorm:"column:payment_id;"`
	Status       string        `json:"status" gorm:"column:status;default:UPCOMING;"`
	CheckInDate  string        `json:"check_in_date" gorm:"column:check_in_date;"`
	CheckOutDate string        `json:"check_out_date" gorm:"column:check_out_date;"`
	Room         RoomJSON      `json:"room" gorm:"type:json;column:room;"`
	User         UserEntity    `json:"user" gorm:"foreignKey:UserID;references:ID;"`
	Payment      PaymentEntity `json:"payment" gorm:"foreignKey:PaymentID;references:ID;"`
}

func (ReservationEntity) TableName() string {
	return "reservations"
}
