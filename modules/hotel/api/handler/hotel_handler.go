package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"trekkstay/modules/hotel/domain/usecase"
)

type HotelHandler interface {
	HandleCreatHotel(c *gin.Context)
}

type hotelHandler struct {
	requestValidator   *validator.Validate
	createHotelUseCase usecase.CreateHotelUseCase
}

func NewHotelHandler(
	requestValidator *validator.Validate,
	createHotelUseCase usecase.CreateHotelUseCase,
) HotelHandler {
	return &hotelHandler{
		requestValidator:   requestValidator,
		createHotelUseCase: createHotelUseCase,
	}
}
