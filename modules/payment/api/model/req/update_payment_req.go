package req

type UpdatePaymentReq struct {
	ReservationID string `json:"reservation_id"`
	Status        string `json:"status" example:"SUCCESS"`
}
