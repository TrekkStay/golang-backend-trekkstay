package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	res "trekkstay/core/response"
	"trekkstay/modules/hotel_room/api/mapper"
	"trekkstay/modules/hotel_room/api/model/req"
	"trekkstay/pkgs/log"
)

// HandleFilterHotelRoom	godoc
// @Summary      Filter hotel room
// @Description  Filter hotel room in a hotel
// @Tags         Hotel Room
// @Produce      json
// @Param        FilterHotelRoomReq  query	req.FilterHotelRoomReq  true  "FilterHotelRoomReq JSON"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /hotel-room/filter [get]
// @Security     JWT
func (h hotelRoomHandler) HandleFilterHotelRoom(c *gin.Context) {
	var req req.FilterHotelRoomReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.JsonLogger.Error("HandleFilterHotelRoom.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	filter := mapper.ConvertFindHotelRoomReqToEntity(req)
	hotelRooms, err := h.filterHotelRoomUseCase.ExecuteFilterHotelRoom(c.Request.Context(), filter)
	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", hotelRooms))
}
