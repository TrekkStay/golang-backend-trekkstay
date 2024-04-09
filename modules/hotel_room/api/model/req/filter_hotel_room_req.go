package req

type FilterHotelRoomReq struct {
	HotelID    *string `form:"hotel_id" extensions:"x-order=1"`
	Balcony    *bool   `form:"balcony" extensions:"x-order=2"`
	BathTub    *bool   `form:"bath_tub" extensions:"x-order=3"`
	Kitchen    *bool   `form:"kitchen" extensions:"x-order=4"`
	NonSmoking *bool   `form:"non_smoking" extensions:"x-order=5"`
	PriceOrder *string `form:"price_order" example:"acs | desc" extensions:"x-order=6"`
}
