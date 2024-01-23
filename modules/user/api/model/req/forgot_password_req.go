package req

type ForgotPasswordReq struct {
	Email string `json:"email" validate:"required,email"`
}
