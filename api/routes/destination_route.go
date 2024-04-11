package routes

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/modules/destination/api/handler"
	"trekkstay/modules/destination/domain/usecase"
	"trekkstay/modules/destination/repository"
	database "trekkstay/pkgs/dbs/postgres"
	"trekkstay/pkgs/dbs/redis"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

func NewDestinationHandler(db *database.Database, requestValidator *validator.Validate) handler.DestinationHandler {
	// Destination Repository
	destinationRepoReader := repository.NewDestinationReaderRepository(*db)
	destinationRepoWriter := repository.NewDestinationWriterRepository(*db)

	// Redis
	redisConfig := config.LoadConfig(&models.RedisConfig{}).(*models.RedisConfig)
	var conn = redis.Connection{
		Address:  fmt.Sprint(redisConfig.RedisHost, ":", redisConfig.RedisPort),
		Password: redisConfig.RedisPassword,
		Database: redisConfig.RedisDB,
	}

	var redisInstance = redis.NewRedis(conn)

	return handler.NewDestinationHandler(
		requestValidator,
		redisInstance,
		usecase.NewCreateDestinationUseCase(destinationRepoWriter),
		usecase.NewListDestinationUseCase(destinationRepoReader),
	)
}

func (r *RouteHandler) destinationRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/destination",
		Routes: []route.Route{
			{
				Path:    "/create",
				Method:  method.POST,
				Handler: r.DestinationHandler.HandleCreateDestination,
			},
			{
				Path:    "/list",
				Method:  method.GET,
				Handler: r.DestinationHandler.HandleListDestination,
			},
		},
	}
}
