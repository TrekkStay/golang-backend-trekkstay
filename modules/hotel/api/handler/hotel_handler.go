package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"trekkstay/modules/hotel/domain/usecase"
)

type HotelHandler interface {
	HandleCreatHotel(c *gin.Context)
	HandleFilterHotel(c *gin.Context)
}

type hotelHandler struct {
	requestValidator   *validator.Validate
	createHotelUseCase usecase.CreateHotelUseCase
	filterHotelUseCase usecase.FilterHotelUseCase
}

func NewHotelHandler(
	requestValidator *validator.Validate,
	createHotelUseCase usecase.CreateHotelUseCase,
	filterHotelUseCase usecase.FilterHotelUseCase,
) HotelHandler {
	return &hotelHandler{
		requestValidator:   requestValidator,
		createHotelUseCase: createHotelUseCase,
		filterHotelUseCase: filterHotelUseCase,
	}
}
