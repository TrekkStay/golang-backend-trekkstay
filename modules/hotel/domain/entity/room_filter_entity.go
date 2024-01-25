package entity

type RoomFilterEntity struct {
	HotelID    *string
	Balcony    *bool
	BathTub    *bool
	Kitchen    *bool
	NonSmoking *bool
}
