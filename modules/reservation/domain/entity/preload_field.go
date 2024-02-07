package entity

type PaymentEntity struct {
	ID            string `json:"id" gorm:"column:id;"`
	ReservationID string `json:"reservation_id" gorm:"column:reservation_id;"`
	Amount        int    `json:"amount" gorm:"column:amount;"`
	Method        string `json:"method" gorm:"column:method;"`
	Status        string `json:"status" gorm:"column:status;"`
}

type UserEntity struct {
	ID       string `json:"id" gorm:"column:id;"`
	FullName string `json:"full_name" gorm:"not null;"`
	Email    string `json:"email" gorm:"uniqueIndex;not null;"`
	Phone    string `json:"phone" gorm:"uniqueIndex;default:null"`
}

func (PaymentEntity) TableName() string {
	return "payments"
}

func (UserEntity) TableName() string {
	return "users"
}
