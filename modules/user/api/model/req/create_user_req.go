package req

type CreateUserReq struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"phone"`
	Password string `json:"password" validate:"required,password"`
}