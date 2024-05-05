package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	res "trekkstay/core/response"
)

// HandleFilterRating	godoc
// @Summary      Filter rating by hotel id
// @Description  Filter rating by hotel id
// @Tags         Rating
// @Produce      json
// @Param        hotel_id  query  string  true  "Hotel ID"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /rating/filter [get]
// @Security     JWT
func (h ratingHandler) HandleFilterRating(c *gin.Context) {
	// Get hotel id from query
	hotelID := c.Query("hotel_id")

	// Call service
	ratings, err := h.filterRatingUseCase.ExecuteGetRatingByHotel(c.Request.Context(), hotelID)

	if err != nil {
		panic(err)
	}

	// Response
	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", ratings))
}
