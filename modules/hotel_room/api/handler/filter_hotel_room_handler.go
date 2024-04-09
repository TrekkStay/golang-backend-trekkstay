package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
	res "trekkstay/core/response"
	"trekkstay/modules/hotel_room/api/mapper"
	"trekkstay/modules/hotel_room/api/model/req"
	"trekkstay/modules/hotel_room/domain/entity"
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

	var hotelRooms []entity.HotelRoomEntity
	cacheKey := c.Request.URL.RequestURI()
	err := h.cache.Get(cacheKey, &hotelRooms)
	if err == nil {
		log.JsonLogger.Debug("cache hit", slog.String("key", cacheKey))
		res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", hotelRooms))
		return
	}

	filter := mapper.ConvertFindHotelRoomReqToEntity(req)
	hotelRooms, err = h.filterHotelRoomUseCase.ExecuteFilterHotelRoom(c.Request.Context(), filter)
	if err != nil {
		panic(err)
	}

	_ = h.cache.SetWithExpiration(cacheKey, hotelRooms, 1*time.Minute)

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", hotelRooms))
}
