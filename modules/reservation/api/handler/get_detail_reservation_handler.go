package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	res "trekkstay/core/response"
)

// HandleGetDetailReservation	godoc
// @Summary      Get detail reservation
// @Description  Get detail reservation
// @Tags         Reservation
// @Produce      json
// @Param        reservation_id  path  string  true  "Reservation ID"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /reservation/{reservation_id} [get]
// @Security     JWT
func (h reservationHandler) HandleGetDetailReservation(c *gin.Context) {
	// Get reservation ID from path
	reservationID := c.Param("reservation_id")

	// Get reservation detail
	reservation, err := h.getDetailReservationUseCase.ExecuteGetDetailReservation(c.Request.Context(), reservationID)

	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", reservation))
}
