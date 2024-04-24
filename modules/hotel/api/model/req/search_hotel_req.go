package req

import "trekkstay/core"

type SearchHotelReq struct {
	core.BaseFilter
	LocationCode *string `form:"location_code" extensions:"x-order=1"`
	PriceOrder   *string `form:"price_order" extensions:"x-order=2"`
	CheckInDate  *string `form:"check_in_date" extensions:"x-order=3" example:"2024-04-24"`
	CheckOutDate *string `form:"check_out_date" extensions:"x-order=4" example:"2024-04-25"`
	Adults       *int    `form:"adults" extensions:"x-order=5" example:"2"`
	Children     *int    `form:"children" extensions:"x-order=6" example:"1"`
	NumOfRooms   *int    `form:"num_of_rooms" extensions:"x-order=7" example:"1"`
}
