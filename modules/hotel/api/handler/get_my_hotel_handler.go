package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"trekkstay/core"
	res "trekkstay/core/response"
)

// HandleGetMyHotel	godoc
// @Summary      Get detail my hotel
// @Description  Get detail my hotel
// @Tags         Hotel
// @Produce      json
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /hotel/my-hotel [get]
// @Security     JWT
func (h hotelHandler) HandleGetMyHotel(c *gin.Context) {
	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	// Get my hotel
	hotel, err := h.getMyHotelUseCase.ExecuteGetMyHotel(ctx)
	if err != nil {
		panic(err)
		return
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", hotel))
}
