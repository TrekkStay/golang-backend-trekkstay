package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"trekkstay/core"
	res "trekkstay/core/response"
	"trekkstay/modules/payment/api/mapper"
	"trekkstay/modules/payment/api/model/req"
	"trekkstay/pkgs/log"
)

// HandleCreatePayment	godoc
// @Summary      Create new payment
// @Description  Create new payment
// @Tags         Payment
// @Produce      json
// @Param        CreatePaymentReq  body	req.CreatePaymentReq  true  "CreatePaymentReq JSON"
// @Success      201 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /payment/create [post]
// @Security     JWT
func (h paymentHandler) HandleCreatePayment(c *gin.Context) {
	// Bind request
	var createPaymentRequest req.CreatePaymentReq
	if err := c.ShouldBindJSON(&createPaymentRequest); err != nil {
		log.JsonLogger.Error("HandleCreatePayment.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	fmt.Print("createPaymentRequest: ", createPaymentRequest)

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	// Create payment
	payment, err := h.createPaymentUseCase.ExecuteCreatePaymentUseCase(
		ctx,
		mapper.ConvertCreatePaymentReqToEntity(createPaymentRequest),
	)

	if err != nil {
		panic(err)
	}

	// Response
	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusCreated, "success", payment))
}
