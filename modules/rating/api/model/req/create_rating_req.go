package req

type CreateRatingReq struct {
	HotelID        string `json:"hotel_id" validate:"required" extensions:"x-order=1"`
	Title          string `json:"title" extensions:"x-order=2"`
	TypeOfTraveler string `json:"type_of_traveler" extensions:"x-order=3"`
	Point          int    `json:"point" extensions:"x-order=4"`
	Summary        string `json:"summary" extensions:"x-order=5"`
}
