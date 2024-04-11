package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"trekkstay/modules/destination/domain/usecase"
	"trekkstay/pkgs/dbs/redis"
)

type DestinationHandler interface {
	HandleCreateDestination(c *gin.Context)
	HandleListDestination(c *gin.Context)
}

type destinationHandler struct {
	requestValidator         *validator.Validate
	cache                    redis.Redis
	createDestinationUseCase usecase.CreateDestinationUseCase
	listDestinationUseCase   usecase.ListDestinationUseCase
}

func NewDestinationHandler(
	requestValidator *validator.Validate,
	cache redis.Redis,
	createDestinationUseCase usecase.CreateDestinationUseCase,
	listDestinationUseCase usecase.ListDestinationUseCase,
) DestinationHandler {
	return &destinationHandler{
		requestValidator:         requestValidator,
		cache:                    cache,
		createDestinationUseCase: createDestinationUseCase,
		listDestinationUseCase:   listDestinationUseCase,
	}
}
