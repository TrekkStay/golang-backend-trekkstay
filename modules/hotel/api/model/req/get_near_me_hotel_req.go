package req

type GetNearMeHotelReq struct {
	Lat         float64 `json:"lat" form:"lat" binding:"required"`
	Lng         float64 `json:"lng" form:"lng" binding:"required"`
	MaxDistance float64 `json:"max_distance" form:"max_distance" binding:"required"`
}
