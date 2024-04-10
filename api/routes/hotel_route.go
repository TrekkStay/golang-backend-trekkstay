package routes

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"trekkstay/api/middlewares"
	"trekkstay/config"
	"trekkstay/config/models"
	hotelHandler "trekkstay/modules/hotel/api/handler"
	"trekkstay/modules/hotel/domain/usecase"
	"trekkstay/modules/hotel/repository"
	database "trekkstay/pkgs/dbs/postgres"
	"trekkstay/pkgs/dbs/redis"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

func NewHotelHandler(db *database.Database, requestValidator *validator.Validate) hotelHandler.HotelHandler {
	// Hotel Repository
	hotelRepoReader := repository.NewHotelReaderRepository(*db)
	hotelRepoWriter := repository.NewHotelWriterRepository(*db)

	// Redis
	redisConfig := config.LoadConfig(&models.RedisConfig{}).(*models.RedisConfig)
	var conn = redis.Connection{
		Address:  fmt.Sprint(redisConfig.RedisHost, ":", redisConfig.RedisPort),
		Password: redisConfig.RedisPassword,
		Database: redisConfig.RedisDB,
	}

	var redisInstance = redis.NewRedis(conn)

	return hotelHandler.NewHotelHandler(requestValidator, redisInstance,
		usecase.NewCreateHotelUseCase(hotelRepoReader, hotelRepoWriter),
		usecase.NewFilterHotelUseCase(hotelRepoReader),
	)
}

func (r *RouteHandler) hotelRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/hotel",
		Routes: []route.Route{
			{
				Path:    "/create",
				Method:  method.POST,
				Handler: r.HotelHandler.HandleCreatHotel,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/filter",
				Method:  method.GET,
				Handler: r.HotelHandler.HandleFilterHotel,
			},
		},
	}
}
