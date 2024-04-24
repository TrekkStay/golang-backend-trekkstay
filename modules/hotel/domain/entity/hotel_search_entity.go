package entity

type HotelSearchEntity struct {
	LocationCode *string
	PriceOrder   *string
	CheckInDate  *string
	CheckOutDate *string
	Adults       *int
	Children     *int
	NumOfRooms   *int
}
