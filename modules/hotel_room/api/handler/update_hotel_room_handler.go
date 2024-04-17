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

// HandleUpdateHotelRoom	godoc
// @Summary      Update hotel room
// @Description  Update hotel room, requires authentication with owner role
// @Tags         Hotel Room
// @Produce      json
// @Param        UpdateHotelRoomReq  body	req.UpdateHotelRoomReq  true  "UpdateHotelRoomReq JSON"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /hotel-room/update [patch]
// @Security     JWT
func (h hotelRoomHandler) HandleUpdateHotelRoom(c *gin.Context) {
	// Bind request
	var updateHotelRoomReq req.UpdateHotelRoomReq
	if err := c.ShouldBind(&updateHotelRoomReq); err != nil {
		log.JsonLogger.Error("HandleUpdateHotelRoom.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Validate request
	if err := h.requestValidator.Struct(&updateHotelRoomReq); err != nil {
		log.JsonLogger.Error("HandleUpdateHotelRoom.validate_request",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrFieldValidationFailed(err))
	}

	// Update hotel room
	if err := h.updateHotelRoomUseCase.ExecUpdateHotelRoomUseCase(
		c.Request.Context(),
		mapper.ConvertUpdateHotelRoomReqToEntity(updateHotelRoomReq),
	); err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", nil))
}
