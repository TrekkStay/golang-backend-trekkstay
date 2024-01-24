package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"trekkstay/core"
	res "trekkstay/core/response"
	"trekkstay/modules/user/api/model/req"
	"trekkstay/pkgs/log"
)

// HandleChangePassword HandleCreateUser	godoc
// @Summary      Change password
// @Description  Change password
// @Tags         User
// @Produce      json
// @Param        ChangePasswordReq  body	req.ChangePasswordReq  true  "ChangePasswordReq JSON"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /user/change-password [post]
// @Security 	jwt
func (h *userHandler) HandleChangePassword(c *gin.Context) {
	// Bind request
	var changePasswordReq req.ChangePasswordReq
	if err := c.ShouldBind(&changePasswordReq); err != nil {
		log.JsonLogger.Error("HandleCreateUser.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Validate request
	if err := h.requestValidator.Struct(&changePasswordReq); err != nil {
		log.JsonLogger.Error("HandleChangePassword.validate_request",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, e := range validationErrors {
				if e.Field() == "OldPwd" {
					panic(res.ErrFieldValidationFailed(errors.New("missing old password")))
				}

				if e.Field() == "NewPwd" {
					panic(res.ErrFieldValidationFailed(errors.New("missing new password")))
				}
			}

			// If no field matched, return default error
			panic(res.ErrFieldValidationFailed(err))
		}
	}

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	// Change password
	if err := h.changePasswordUseCase.ExecChangePassword(ctx,
		changePasswordReq.OldPwd, changePasswordReq.NewPwd); err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "change password successfully", nil))
}
