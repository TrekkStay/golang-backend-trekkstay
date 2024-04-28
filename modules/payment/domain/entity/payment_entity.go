package entity

import "trekkstay/core"

type PaymentEntity struct {
	core.BaseEntity
	ReservationID string `json:"reservation_id" gorm:"column:reservation_id;"`
	ReturnCode    string `json:"return_code" gorm:"column:return_code;default:null"`
	Amount        int    `json:"amount" gorm:"column:amount;"`
	Method        string `json:"method" gorm:"column:method;"`
	Status        string `json:"status" gorm:"column:status;default:PENDING"`
}
