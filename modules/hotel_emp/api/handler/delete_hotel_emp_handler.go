package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	res "trekkstay/core/response"
)

// HandleDeleteHotelEmp	godoc
// @Summary      Delete hotel employee
// @Description  Delete hotel employee, requires authentication with owner role
// @Tags         Hotel Employee
// @Produce      json
// @Param        employee_id path string true "employee_id"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /hotel-emp/{employee_id} [delete]
// @Security     JWT
func (h hotelEmpHandler) HandleDeleteHotelEmp(c *gin.Context) {
	employeeID := c.Param("employee_id")

	if err := h.deleteHotelEmpUseCase.ExecuteDeleteHotelEmp(c, employeeID); err != nil {
		panic(res.ErrInvalidRequest(err))
		return
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", nil))
}
