package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	res "trekkstay/core/response"
)

// HandleGetDetailHotelRoom	godoc
// @Summary      Get detail hotel's room
// @Description  Get detail hotel's room
// @Tags         Hotel Room
// @Produce      json
// @Param        room_id  path  string  true  "Room ID"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /hotel-room/{hotel_room_id} [get]
func (h hotelRoomHandler) HandleGetDetailHotelRoom(c *gin.Context) {
	hotelRoomID := c.Param("hotel_room_id")

	if hotelRoomID == "" {
		panic(res.ErrInvalidRequest(errors.New("room's id is required")))
		return
	}

	hotelRoom, err := h.getDetailHotelRoomUseCase.ExecuteGetDetailHotelRoom(c.Request.Context(), hotelRoomID)
	if err != nil {
		panic(err)
		return
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", hotelRoom))
}
