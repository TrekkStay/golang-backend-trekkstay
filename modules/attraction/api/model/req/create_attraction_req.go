package req

type CreateAttractionReq struct {
	Name         string  `json:"name" validate:"required" extensions:"x-order=1"`
	Lat          float64 `json:"lat" validate:"required" extensions:"x-order=2"`
	Lng          float64 `json:"lng" validate:"required" extensions:"x-order=3"`
	ProvinceCode string  `json:"province_code" validate:"required" extensions:"x-order=4"`
	DistrictCode string  `json:"district_code" validate:"required" extensions:"x-order=5"`
	WardCode     string  `json:"ward_code" validate:"required" extensions:"x-order=6"`
}
