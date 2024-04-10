package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	res "trekkstay/core/response"
)

// HandleGetDetailHotel	godoc
// @Summary      Get detail hotel
// @Description  Get detail hotel
// @Tags         Hotel
// @Produce      json
// @Param        hotel_id  path  string  true  "Hotel ID"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /hotel/{hotel_id} [post]
// @Security     JWT
func (h hotelHandler) HandleGetDetailHotel(c *gin.Context) {
	hotelID := c.Param("hotel_id")

	if hotelID == "" {
		panic(res.ErrInvalidRequest(errors.New("hotel's id is required")))
		return
	}

	hotel, err := h.getDetailHotelUseCase.ExecuteGetDetailHotel(c.Request.Context(), hotelID)
	if err != nil {
		panic(err)
		return
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", hotel))
}
