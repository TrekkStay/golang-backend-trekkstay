package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"trekkstay/modules/rating/domain/usecase"
)

type RatingHandler interface {
	HandleCreateRating(c *gin.Context)
	HandleFilterRating(c *gin.Context)
}

type ratingHandler struct {
	requestValidator    *validator.Validate
	createRatingUseCase usecase.CreateRatingUseCase
	filterRatingUseCase usecase.GetRatingByHotelUseCase
}

func NewRatingHandler(
	requestValidator *validator.Validate,
	createRatingUseCase usecase.CreateRatingUseCase,
	filterRatingUseCase usecase.GetRatingByHotelUseCase,
) RatingHandler {
	return &ratingHandler{
		requestValidator:    requestValidator,
		createRatingUseCase: createRatingUseCase,
		filterRatingUseCase: filterRatingUseCase,
	}
}
