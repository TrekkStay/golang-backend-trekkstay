package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	res "trekkstay/core/response"
)

// HandleFilterHotelEmp	godoc
// @Summary      Filter hotel employee
// @Description  Filter hotel employee
// @Tags         Hotel Employee
// @Produce      json
// @Param        hotel_id  query  string  true  "Hotel ID"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /hotel-emp/filter [get]
// @Security 	JWT
func (h hotelEmpHandler) HandleFilterHotelEmp(c *gin.Context) {
	hotelID := c.Query("hotel_id")

	hotelEmployees, err := h.filterHotelEmpUseCase.ExecuteFilterHotelEmp(c.Request.Context(), hotelID)

	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", hotelEmployees))
}
