package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"trekkstay/modules/hotel_room/domain/usecase"
)

type HotelRoomHandler interface {
	HandleCreateHotelRoom(c *gin.Context)
}

type hotelRoomHandler struct {
	requestValidator       *validator.Validate
	createHotelRoomUseCase usecase.CreateHotelRoomUseCase
}

func NewHotelRoomHandler(
	requestValidator *validator.Validate,
	createHotelRoomUseCase usecase.CreateHotelRoomUseCase,
) HotelRoomHandler {
	return &hotelRoomHandler{
		requestValidator:       requestValidator,
		createHotelRoomUseCase: createHotelRoomUseCase,
	}
}
