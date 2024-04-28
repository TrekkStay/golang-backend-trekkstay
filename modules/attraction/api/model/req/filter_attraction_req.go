package req

type FilterAttractionReq struct {
	LocationCode string `json:"location_code" form:"location_code" validate:"required"`
}
