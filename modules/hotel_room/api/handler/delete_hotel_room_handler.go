package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	res "trekkstay/core/response"
)

// HandleDeleteHotelRoom	godoc
// @Summary      Delete hotel room
// @Description  Delete hotel room, requires authentication with owner role
// @Tags         Hotel Room
// @Produce      json
// @Param        room_id path string true "room_id"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /hotel-room/{room_id} [delete]
// @Security     JWT
func (h hotelRoomHandler) HandleDeleteHotelRoom(c *gin.Context) {
	roomID := c.Param("room_id")

	err := h.deleteHotelRoomUseCase.ExecuteDeleteHotelRoom(c.Request.Context(), roomID)

	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", nil))
}
