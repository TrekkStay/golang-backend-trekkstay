package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"trekkstay/modules/hotel_room/domain/usecase"
	"trekkstay/pkgs/dbs/redis"
)

type HotelRoomHandler interface {
	HandleCreateHotelRoom(c *gin.Context)
	HandleFilterHotelRoom(c *gin.Context)
	HandleUpdateHotelRoom(c *gin.Context)
	HandleGetDetailHotelRoom(c *gin.Context)
	HandleDeleteHotelRoom(c *gin.Context)
}

type hotelRoomHandler struct {
	requestValidator          *validator.Validate
	cache                     redis.Redis
	createHotelRoomUseCase    usecase.CreateHotelRoomUseCase
	filterHotelRoomUseCase    usecase.FilterHotelRoomUseCase
	updateHotelRoomUseCase    usecase.UpdateHotelRoomUseCase
	getDetailHotelRoomUseCase usecase.GetDetailHotelRoomUseCase
	deleteHotelRoomUseCase    usecase.DeleteHotelRoomUseCase
}

func NewHotelRoomHandler(
	requestValidator *validator.Validate,
	cache redis.Redis,
	createHotelRoomUseCase usecase.CreateHotelRoomUseCase,
	filterHotelRoomUseCase usecase.FilterHotelRoomUseCase,
	updateHotelRoomUseCase usecase.UpdateHotelRoomUseCase,
	getDetailHotelRoomUseCase usecase.GetDetailHotelRoomUseCase,
	deleteHotelRoomUseCase usecase.DeleteHotelRoomUseCase,
) HotelRoomHandler {
	return &hotelRoomHandler{
		requestValidator:          requestValidator,
		cache:                     cache,
		createHotelRoomUseCase:    createHotelRoomUseCase,
		filterHotelRoomUseCase:    filterHotelRoomUseCase,
		updateHotelRoomUseCase:    updateHotelRoomUseCase,
		getDetailHotelRoomUseCase: getDetailHotelRoomUseCase,
		deleteHotelRoomUseCase:    deleteHotelRoomUseCase,
	}
}
