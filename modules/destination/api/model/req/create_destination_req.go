package req

type CreateDestinationReq struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
	Unit string `json:"unit" validate:"required"`
}
