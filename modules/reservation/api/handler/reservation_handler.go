package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"trekkstay/modules/reservation/domain/usecase"
)

type ReservationHandler interface {
	HandleCreateReservation(c *gin.Context)
	HandleFilterReservation(c *gin.Context)
}

type reservationHandler struct {
	requestValidator         *validator.Validate
	createReservationUseCase usecase.CreateReservationUseCase
	filterReservationUseCase usecase.FilterReservationUseCase
}

func NewReservationHandler(
	requestValidator *validator.Validate,
	createReservationUseCase usecase.CreateReservationUseCase,
	filterReservationUseCase usecase.FilterReservationUseCase,
) ReservationHandler {
	return &reservationHandler{
		requestValidator:         requestValidator,
		createReservationUseCase: createReservationUseCase,
		filterReservationUseCase: filterReservationUseCase,
	}
}
