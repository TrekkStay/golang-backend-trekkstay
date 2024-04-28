package req

type GuestInfo struct {
	FullName string `json:"full_name"`
	Contact  string `json:"contact"`
	Adults   int    `json:"adults"`
	Children int    `json:"children"`
}

type CreateReservationReq struct {
	RoomID        string    `json:"room_id" validate:"required" extensions:"x-order=1"`
	PromotionCode string    `json:"promotion_code" extensions:"x-order=2"`
	CheckInDate   string    `json:"check_in_date" validate:"required" extensions:"x-order=3"`
	CheckOutDate  string    `json:"check_out_date" validate:"required" extensions:"x-order=4"`
	Quantity      int       `json:"quantity" validate:"required" extensions:"x-order=5"`
	Guest         GuestInfo `json:"guest_info" validate:"required" extensions:"x-order=6"`
}
