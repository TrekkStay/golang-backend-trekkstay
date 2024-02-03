package req

type LoginHotelEmpReq struct {
	Email    string `json:"email" validate:"required,email" extensions:"x-order=1"`
	Password string `json:"password" validate:"required,password" extensions:"x-order=2"`
}
