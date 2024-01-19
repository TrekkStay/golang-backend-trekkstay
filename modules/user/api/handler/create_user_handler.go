package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	res "trekkstay/core/response"
	"trekkstay/modules/user/api/mapper"
	"trekkstay/modules/user/api/model/req"
	"trekkstay/pkgs/log"
)

// HandleCreateUser	godoc
// @Summary      Register new user
// @Description  Register new user
// @Tags         User
// @Produce      json
// @Param        CreateUserReq  body	req.CreateUserReq  true  "CreateUserReq JSON"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /api/v1/user/signup [post]
// @Security     JWT
func (h *userHandler) HandleCreateUser(c *gin.Context) {
	// Bind request
	var createUserReq req.CreateUserReq
	if err := c.ShouldBindJSON(&createUserReq); err != nil {
		log.JsonLogger.Error("HandleCreateUser.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Validate request
	if err := h.requestValidator.Struct(&createUserReq); err != nil {
		log.JsonLogger.Error("HandleCreateUser.validate_request",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrFieldValidationFailed(err))
	}

	// Create user
	if err := h.createUserUseCase.ExecCreateUser(c.Request.Context(),
		mapper.ConvertCreateUserReqToUserEntity(createUserReq)); err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", nil))
}
