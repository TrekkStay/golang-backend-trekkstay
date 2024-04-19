package req

type CreateAttractionReq struct {
	Name         string `json:"name" validate:"required"`
	ProvinceCode string `json:"province_code" validate:"required"`
	DistrictCode string `json:"district_code" validate:"required"`
	WardCode     string `json:"ward_code" validate:"required"`
}
