package req

import "trekkstay/core"

type FilterHotelReq struct {
	core.BaseFilter
	Name         *string `form:"name" extensions:"x-order=1"`
	ProvinceCode *string `form:"province_code" extensions:"x-order=2"`
	DistrictCode *string `form:"district_code" extensions:"x-order=3"`
	WardCode     *string `form:"ward_code" extensions:"x-order=4"`
	PriceOrder   *string `form:"price_order"  example:"acs | desc" extensions:"x-order=5"`
}
