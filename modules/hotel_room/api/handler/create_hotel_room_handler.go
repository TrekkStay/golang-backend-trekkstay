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

// HandleCreateHotelRoom	godoc
// @Summary      Create new hotel room
// @Description  Create new hotel room, requires authentication with owner role
// @Tags         Hotel Room
// @Produce      json
// @Param        CreateHotelRoomReq  body	req.CreateHotelRoomReq  true  "CreateHotelRoomReq JSON"
// @Success      201 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /hotel-room/create [post]
// @Security     JWT
func (h hotelRoomHandler) HandleCreateHotelRoom(c *gin.Context) {
	// Bind request
	var createHotelRoomReq req.CreateHotelRoomReq
	if err := c.ShouldBindJSON(&createHotelRoomReq); err != nil {
		log.JsonLogger.Error("HandleCreateHotelRoom.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Validate request
	if err := h.requestValidator.Struct(&createHotelRoomReq); err != nil {
		log.JsonLogger.Error("HandleCreateHotelRoom.validate_request",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrFieldValidationFailed(err))
	}

	hotel := mapper.ConvertCreateHotelRoomReqToEntity(createHotelRoomReq)

	// Create hotel room
	if err := h.createHotelRoomUseCase.ExecuteCreateHotelRoom(
		c.Request.Context(),
		&hotel,
	); err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusCreated, "success", hotel.HotelID))
}
