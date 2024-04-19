package routes

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/modules/attraction/api/handler"
	"trekkstay/modules/attraction/domain/usecase"
	"trekkstay/modules/attraction/repository"
	database "trekkstay/pkgs/dbs/postgres"
	"trekkstay/pkgs/dbs/redis"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

func NewAttractionHandler(db *database.Database, requestValidator *validator.Validate) handler.AttractionHandler {
	// Destination Repository
	attractionRepoReader := repository.NewAttractionReaderRepository(*db)
	attractionRepoWriter := repository.NewAttractionWriterRepository(*db)

	// Redis
	redisConfig := config.LoadConfig(&models.RedisConfig{}).(*models.RedisConfig)
	var conn = redis.Connection{
		Address:  fmt.Sprint(redisConfig.RedisHost, ":", redisConfig.RedisPort),
		Password: redisConfig.RedisPassword,
		Database: redisConfig.RedisDB,
	}

	var redisInstance = redis.NewRedis(conn)

	return handler.NewAttractionHandler(
		requestValidator,
		redisInstance,
		usecase.NewCreateAttractionUseCase(attractionRepoWriter),
		usecase.NewListAttractionUseCase(attractionRepoReader),
	)
}

func (r *RouteHandler) attractionRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/attraction",
		Routes: []route.Route{
			{
				Path:    "/create",
				Method:  method.POST,
				Handler: r.AttractionHandler.HandleCreateAttraction,
			},
			{
				Path:    "/list",
				Method:  method.GET,
				Handler: r.AttractionHandler.HandleListAttraction,
			},
		},
	}
}
