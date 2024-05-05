package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	res "trekkstay/core/response"
)

// HandleCancelReservation	godoc
// @Summary      Cancel reservation
// @Description  Cancel reservation
// @Tags         Reservation
// @Produce      json
// @Param        reservation_id  path  string  true  "Reservation ID"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /reservation/cancel/{reservation_id} [post]
// @Security     JWT
func (h reservationHandler) HandleCancelReservation(c *gin.Context) {
	reservationID := c.Param("reservation_id")
	err := h.cancelReservationUseCase.ExecuteCancelReservation(c, reservationID)
	if err != nil {
		panic(err)
		return
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", nil))
}
