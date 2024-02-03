package res

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginHotelEmpRes struct {
	HotelID  string `json:"hotel_id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Token    Token  `json:"token"`
}
