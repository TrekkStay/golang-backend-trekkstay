package entity

type HotelRoomFilterEntity struct {
	HotelID    *string
	Balcony    *bool
	BathTub    *bool
	Kitchen    *bool
	NonSmoking *bool
	PriceOrder *string
}
