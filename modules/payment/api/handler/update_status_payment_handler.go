package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"trekkstay/core"
	res "trekkstay/core/response"
	"trekkstay/modules/payment/api/model/req"
	"trekkstay/pkgs/log"
)

// HandleUpdatePayment	godoc
// @Summary      Update payment
// @Description  Update payment
// @Tags         Payment
// @Produce      json
// @Param        UpdatePaymentReq  body	req.UpdatePaymentReq  true  "UpdatePaymentReq JSON"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /payment/update [patch]
// @Security     JWT
func (h paymentHandler) HandleUpdatePayment(c *gin.Context) {
	// Bind request
	var updatePaymentReq req.UpdatePaymentReq
	if err := c.ShouldBindJSON(&updatePaymentReq); err != nil {
		log.JsonLogger.Error("HandleUpdatePayment.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	// Update payment status
	err := h.updatePaymentUseCase.ExecuteUpdateStatusPayment(
		ctx,
		updatePaymentReq.ReservationID,
		updatePaymentReq.Status,
	)

	if err != nil {
		panic(err)
	}

	// Response
	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", nil))
}
