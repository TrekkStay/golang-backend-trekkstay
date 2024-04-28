package req

import "trekkstay/core"

type FilterReservationReq struct {
	core.BaseFilter
	HotelID *string `form:"hotel_id" extensions:"x-order=1"`
	Status  *string `form:"status" extensions:"x-order=2" example:"UPCOMING | COMPLETED | CANCELLED"`
	Date    *string `form:"date" extensions:"x-order=3"`
}
