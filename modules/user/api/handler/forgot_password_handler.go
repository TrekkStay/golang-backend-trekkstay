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

// HandleForgotPassword HandleChangePassword HandleCreateUser	godoc
// @Summary      Forgot password
// @Description  Forgot password and send new password to email
// @Tags         User
// @Produce      json
// @Param        ForgotPasswordReq  body	req.ForgotPasswordReq  true  "ForgotPasswordReq JSON"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /user/forgot-password [post]
func (h *userHandler) HandleForgotPassword(c *gin.Context) {
	// Bind request
	var forgotPasswordReq req.ForgotPasswordReq
	if err := c.ShouldBind(&forgotPasswordReq); err != nil {
		log.JsonLogger.Error("HandleForgotPassword.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Validate request
	if err := h.requestValidator.Struct(&forgotPasswordReq); err != nil {
		log.JsonLogger.Error("HandleForgotPassword.validate_request",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrFieldValidationFailed(errors.New("missing email")))
	}

	// Execute forgot password use case
	if err := h.forgotPasswordUseCase.ExecuteForgotPassword(c.Request.Context(), forgotPasswordReq.Email); err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "new password was sent to your email", nil))
}
