package req

type ChangePasswordReq struct {
	OldPwd string `json:"old_pwd" validate:"required"`
	NewPwd string `json:"new_pwd" validate:"required"`
}
