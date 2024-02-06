package req

type CreateHotelRoomReq struct {
	HotelID       string                  `json:"hotel_id" validate:"required" extensions:"x-order=1"`
	Type          string                  `json:"type" validate:"required" extensions:"x-order=2"`
	Description   string                  `json:"description" validate:"required" extensions:"x-order=3"`
	Quantity      int                     `json:"quantity" validate:"required" extensions:"x-order=4"`
	DiscountRate  int                     `json:"discount_rate" extensions:"x-order=5"`
	OriginalPrice int                     `json:"original_price" validate:"required" extensions:"x-order=6"`
	Videos        MediaJSON               `json:"videos" extensions:"x-order=7"`
	Images        MediaJSON               `json:"images" extensions:"x-order=8"`
	Facilities    HotelRoomFacilitiesJSON `json:"facilities" validate:"required" extensions:"x-order=9"`
}
