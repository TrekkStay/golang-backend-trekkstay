package req

import "trekkstay/core"

type FilterHotelReq struct {
	core.BaseFilter
	Name         *string `form:"name"`
	ProvinceCode *string `form:"province_code"`
	DistrictCode *string `form:"district_code"`
	WardCode     *string `form:"ward_code"`
	PriceOrder   *string `form:"price_order"  example:"acs | desc"`
}
