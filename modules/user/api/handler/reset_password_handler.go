package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	res "trekkstay/core/response"
	"trekkstay/modules/user/api/model/req"
	"trekkstay/pkgs/log"
)

// HandleResetPassword godoc
// @Summary      Reset password
// @Description  Reset password and send new password to email
// @Tags         User
// @Produce      json
// @Param        ResetPasswordReq  body	req.ResetPasswordReq  true  "ResetPasswordReq JSON"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /user/reset-password [post]
func (h *userHandler) HandleResetPassword(c *gin.Context) {
	// Bind request
	var resetPasswordReq req.ResetPasswordReq
	if err := c.ShouldBind(&resetPasswordReq); err != nil {
		log.JsonLogger.Error("HandleResetPassword.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Validate request
	if err := h.requestValidator.Struct(&resetPasswordReq); err != nil {
		log.JsonLogger.Error("HandleResetPassword.validate_request",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrFieldValidationFailed(errors.New("missing email")))
	}

	// ExecuteGetDetailHotel forgot password use case
	if err := h.resetPasswordUseCase.ExecuteResetPassword(c.Request.Context(), resetPasswordReq.Email); err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "new password was sent to your email", nil))
}
