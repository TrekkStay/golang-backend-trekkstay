package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"trekkstay/core"
	res "trekkstay/core/response"
	"trekkstay/modules/user/api/mapper"
)

// HandleRefreshToken godoc
// @Summary      Refresh token
// @Description  Get new access token and refresh token
// @Tags         User
// @Produce      json
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /user/refresh-token [get]
// @Security     JWT
func (h *userHandler) HandleRefreshToken(c *gin.Context) {
	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	// Execute use case
	accessToken, refreshToken, err := h.refreshTokenUseCase.ExecRefreshToken(ctx)

	if err != nil {
		panic(err)
	}

	// Convert token to response
	token := mapper.ConvertTokenToResponse(*accessToken, *refreshToken)
	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", token))
}
