package entity

import "trekkstay/core"

type PaymentEntity struct {
	core.BaseEntity
	ReservationID string `json:"reservation_id" gorm:"column:reservation_id;"`
	UserID        string `json:"user_id" gorm:"column:user_id;"`
	Amount        int    `json:"amount" gorm:"column:amount;"`
	Method        string `json:"method" gorm:"column:method;default:PAY_AT_HOTEL"`
	Status        string `json:"status" gorm:"column:status;default:PENDING"`
}

func (PaymentEntity) TableName() string {
	return "payments"
}
