package routes

import (
	"github.com/go-playground/validator/v10"
	"trekkstay/api/middlewares"
	"trekkstay/modules/rating/api/handler"
	"trekkstay/modules/rating/domain/usecase"
	"trekkstay/modules/rating/repository"
	database "trekkstay/pkgs/dbs/postgres"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

func NewRatingHandler(db *database.Database, requestValidator *validator.Validate) handler.RatingHandler {
	// Payment Repository
	ratingRepoReader := repository.NewRatingReaderRepository(*db)
	ratingRepoWriter := repository.NewRatingWriterRepository(*db)

	return handler.NewRatingHandler(
		requestValidator,
		usecase.NewCreateRatingUseCase(ratingRepoWriter),
		usecase.NewGetRatingByHotelUseCase(ratingRepoReader),
	)
}

func (r *RouteHandler) ratingRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/rating",
		Routes: []route.Route{
			{
				Path:    "/create",
				Method:  method.POST,
				Handler: r.RatingHandler.HandleCreateRating,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/filter",
				Method:  method.GET,
				Handler: r.RatingHandler.HandleFilterRating,
			},
		},
	}
}
