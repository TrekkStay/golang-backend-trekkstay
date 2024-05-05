package req

type CreatePaymentReq struct {
	ReservationID string `json:"reservation_id"`
	Amount        int    `json:"amount" example:"1000000"`
	Method        string `json:"method" example:"MOMO"`
	Status        string `json:"status" example:"PENDING"`
}
