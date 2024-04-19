package req

type FilterAttractionReq struct {
	ProvinceCode *string `form:"province_code" json:"province_code"`
	DistrictCode *string `form:"district_code" json:"district_code"`
	WardCode     *string `form:"ward_code" json:"ward_code"`
}
