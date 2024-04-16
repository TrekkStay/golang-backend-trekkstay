package req

type UpdateHotelReq struct {
	Name          string              `json:"name" validate:"required" extensions:"x-order=1"`
	Email         string              `json:"email" validate:"required,email" extensions:"x-order=2"`
	Phone         string              `json:"phone" validate:"required" extensions:"x-order=3"`
	CheckInTime   string              `json:"check_in_time" validate:"required" extensions:"x-order=4"`
	CheckOutTime  string              `json:"check_out_time" validate:"required" extensions:"x-order=5"`
	ProvinceCode  string              `json:"province_code" validate:"required" extensions:"x-order=6"`
	DistrictCode  string              `json:"district_code" validate:"required" extensions:"x-order=7"`
	WardCode      string              `json:"ward_code" validate:"required" extensions:"x-order=8"`
	AddressDetail string              `json:"address_detail" validate:"required" extensions:"x-order=9"`
	Description   string              `json:"description" validate:"required" extensions:"x-order=10"`
	Facilities    HotelFacilitiesJSON `json:"facilities" extensions:"x-order=11"`
	Coordinates   CoordinatesJSON     `json:"coordinates" extensions:"x-order=12"`
	Videos        MediaJSON           `json:"videos" extensions:"x-order=13"`
	Images        MediaJSON           `json:"images" extensions:"x-order=14"`
}
