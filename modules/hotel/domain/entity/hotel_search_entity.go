package entity

type HotelSearchEntity struct {
	LocationCode  *string
	AttractionLat *float64
	AttractionLng *float64
	PriceOrder    *string
	CheckInDate   *string
	CheckOutDate  *string
	Adults        *int
	Children      *int
	NumOfRooms    *int
}
