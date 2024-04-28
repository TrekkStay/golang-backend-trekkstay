package handler

import (
	"github.com/gin-gonic/gin"
	"trekkstay/modules/payment/domain/usecase"
)

type PaymentHandler interface {
	HandleCreatePayment(c *gin.Context)
	HandleGetDetailPayment(c *gin.Context)
	HandleUpdatePayment(c *gin.Context)
}

type paymentHandler struct {
	createPaymentUseCase    usecase.CreatePaymentUseCase
	getDetailPaymentUseCase usecase.GetDetailPaymentUseCase
	updatePaymentUseCase    usecase.UpdateStatusPaymentUseCase
}

func NewPaymentHandler(
	createPaymentUseCase usecase.CreatePaymentUseCase,
	getDetailPaymentUseCase usecase.GetDetailPaymentUseCase,
	updatePaymentUseCase usecase.UpdateStatusPaymentUseCase,
) PaymentHandler {
	return &paymentHandler{
		createPaymentUseCase:    createPaymentUseCase,
		getDetailPaymentUseCase: getDetailPaymentUseCase,
		updatePaymentUseCase:    updatePaymentUseCase,
	}
}
