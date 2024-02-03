package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"trekkstay/core"
	res "trekkstay/core/response"
)

// HandleRefreshToken godoc
// @Summary      Refresh token
// @Description  Get new access token and refresh token
// @Tags         Token
// @Produce      json
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /token/refresh-token [get]
// @Security     JWT
func (h *tokenHandler) HandleRefreshToken(c *gin.Context) {
	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	// Execute use case
	token, err := h.refreshTokenUseCase.ExecRefreshToken(ctx)

	if err != nil {
		panic(err)
	}

	// Convert token to response
	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", token))
}
