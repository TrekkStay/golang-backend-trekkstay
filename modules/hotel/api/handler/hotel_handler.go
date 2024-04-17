package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"trekkstay/modules/hotel/domain/usecase"
	"trekkstay/pkgs/dbs/redis"
)

type HotelHandler interface {
	HandleCreatHotel(c *gin.Context)
	HandleFilterHotel(c *gin.Context)
	HandleGetDetailHotel(c *gin.Context)
	HandleGetMyHotel(c *gin.Context)
	HandleUpdateHotel(c *gin.Context)
	HandleGetNearMeHotel(c *gin.Context)
}

type hotelHandler struct {
	requestValidator      *validator.Validate
	cache                 redis.Redis
	createHotelUseCase    usecase.CreateHotelUseCase
	filterHotelUseCase    usecase.FilterHotelUseCase
	getDetailHotelUseCase usecase.GetDetailHotelUseCase
	getMyHotelUseCase     usecase.GetMyHotelUseCase
	updateHotelUseCase    usecase.UpdateHotelUseCase
	getNearMeHotelUseCase usecase.GetNearMeHotelUseCase
}

func NewHotelHandler(
	requestValidator *validator.Validate,
	cache redis.Redis,
	createHotelUseCase usecase.CreateHotelUseCase,
	filterHotelUseCase usecase.FilterHotelUseCase,
	getDetailHotelUseCase usecase.GetDetailHotelUseCase,
	getMyHotelUseCase usecase.GetMyHotelUseCase,
	updateHotelUseCase usecase.UpdateHotelUseCase,
	getNearMeHotelUseCase usecase.GetNearMeHotelUseCase,
) HotelHandler {
	return &hotelHandler{
		requestValidator:      requestValidator,
		cache:                 cache,
		createHotelUseCase:    createHotelUseCase,
		filterHotelUseCase:    filterHotelUseCase,
		getDetailHotelUseCase: getDetailHotelUseCase,
		getMyHotelUseCase:     getMyHotelUseCase,
		updateHotelUseCase:    updateHotelUseCase,
		getNearMeHotelUseCase: getNearMeHotelUseCase,
	}
}
