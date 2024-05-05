package entity

type PaymentEntity struct {
	ReservationID string `json:"reservation_id"`
	UserID        string `json:"user_id"`
	Amount        int    `json:"amount"`
	Method        string `json:"method"`
	Status        string `json:"status"`
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
