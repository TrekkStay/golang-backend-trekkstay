package routes

import (
	"github.com/go-playground/validator/v10"
	"trekkstay/api/middlewares"
	"trekkstay/modules/payment/api/handler"
	"trekkstay/modules/payment/domain/usecase"
	"trekkstay/modules/payment/repository"
	database "trekkstay/pkgs/dbs/postgres"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

func NewPaymentHandler(db *database.Database, requestValidator *validator.Validate) handler.PaymentHandler {
	// Payment Repository
	paymentRepoWriter := repository.NewPaymentWriterRepository(*db)

	// Redis
	return handler.NewPaymentHandler(
		usecase.NewCreatePaymentUseCase(paymentRepoWriter),
		usecase.NewUpdateStatusPaymentUseCase(paymentRepoWriter),
	)
}

func (r *RouteHandler) paymentRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/payment",
		Routes: []route.Route{
			{
				Path:    "/create",
				Method:  method.POST,
				Handler: r.PaymentHandler.HandleCreatePayment,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/update",
				Method:  method.PATCH,
				Handler: r.PaymentHandler.HandleUpdatePayment,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
		},
	}
}
