package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"trekkstay/modules/attraction/domain/usecase"
	"trekkstay/pkgs/dbs/redis"
)

type AttractionHandler interface {
	HandleCreateAttraction(c *gin.Context)
	HandleListAttraction(c *gin.Context)
}

type attractionHandler struct {
	requestValidator        *validator.Validate
	cache                   redis.Redis
	createAttractionUseCase usecase.CreateAttractionUseCase
	listAttractionUseCase   usecase.ListAttractionUseCase
}

func NewAttractionHandler(
	requestValidator *validator.Validate,
	cache redis.Redis,
	createDestinationUseCase usecase.CreateAttractionUseCase,
	listDestinationUseCase usecase.ListAttractionUseCase,
) AttractionHandler {
	return &attractionHandler{
		requestValidator:        requestValidator,
		cache:                   cache,
		createAttractionUseCase: createDestinationUseCase,
		listAttractionUseCase:   listDestinationUseCase,
	}
}
