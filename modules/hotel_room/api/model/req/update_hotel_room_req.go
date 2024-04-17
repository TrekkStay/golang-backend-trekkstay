package req

type UpdateHotelRoomReq struct {
	ID            string                  `json:"id" validate:"required" extensions:"x-order=1"`
	HotelID       string                  `json:"hotel_id" validate:"required" extensions:"x-order=2"`
	Type          string                  `json:"type" validate:"required" extensions:"x-order=3"`
	Description   string                  `json:"description" validate:"required" extensions:"x-order=4"`
	Quantity      int                     `json:"quantity" validate:"required" extensions:"x-order=5"`
	DiscountRate  int                     `json:"discount_rate" extensions:"x-order=6"`
	OriginalPrice int                     `json:"original_price" validate:"required" extensions:"x-order=7"`
	Videos        MediaJSON               `json:"videos" extensions:"x-order=8"`
	Images        MediaJSON               `json:"images" extensions:"x-order=9"`
	Facilities    HotelRoomFacilitiesJSON `json:"facilities" validate:"required" extensions:"x-order=10"`
}
