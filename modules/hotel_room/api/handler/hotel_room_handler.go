package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"trekkstay/modules/hotel_room/domain/usecase"
)

type HotelRoomHandler interface {
	HandleCreateHotelRoom(c *gin.Context)
	HandleFilterHotelRoom(c *gin.Context)
}

type hotelRoomHandler struct {
	requestValidator       *validator.Validate
	createHotelRoomUseCase usecase.CreateHotelRoomUseCase
	filterHotelRoomUseCase usecase.FilterHotelRoomUseCase
}

func NewHotelRoomHandler(
	requestValidator *validator.Validate,
	createHotelRoomUseCase usecase.CreateHotelRoomUseCase,
	filterHotelRoomUseCase usecase.FilterHotelRoomUseCase,
) HotelRoomHandler {
	return &hotelRoomHandler{
		requestValidator:       requestValidator,
		createHotelRoomUseCase: createHotelRoomUseCase,
		filterHotelRoomUseCase: filterHotelRoomUseCase,
	}
}
