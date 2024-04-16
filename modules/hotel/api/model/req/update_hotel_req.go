package req

type UpdateHotelReq struct {
	ID            string              `json:"id" validate:"required" extensions:"x-order=1"`
	Name          string              `json:"name" validate:"required" extensions:"x-order=2"`
	Email         string              `json:"email" validate:"required,email" extensions:"x-order=3"`
	Phone         string              `json:"phone" validate:"required" extensions:"x-order=4"`
	CheckInTime   string              `json:"check_in_time" validate:"required" extensions:"x-order=5"`
	CheckOutTime  string              `json:"check_out_time" validate:"required" extensions:"x-order=6"`
	ProvinceCode  string              `json:"province_code" validate:"required" extensions:"x-order=7"`
	DistrictCode  string              `json:"district_code" validate:"required" extensions:"x-order=8"`
	WardCode      string              `json:"ward_code" validate:"required" extensions:"x-order=9"`
	AddressDetail string              `json:"address_detail" validate:"required" extensions:"x-order=10"`
	Description   string              `json:"description" validate:"required" extensions:"x-order=11"`
	Facilities    HotelFacilitiesJSON `json:"facilities" extensions:"x-order=12"`
	Coordinates   CoordinatesJSON     `json:"coordinates" extensions:"x-order=13"`
	Videos        MediaJSON           `json:"videos" extensions:"x-order=14"`
	Images        MediaJSON           `json:"images" extensions:"x-order=15"`
}
