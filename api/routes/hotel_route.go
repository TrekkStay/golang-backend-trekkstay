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
	emp "trekkstay/modules/hotel_emp/repository"

	database "trekkstay/pkgs/dbs/postgres"
	"trekkstay/pkgs/dbs/redis"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

func NewHotelHandler(db *database.Database, requestValidator *validator.Validate) hotelHandler.HotelHandler {
	// Hotel Repository
	hotelRepoReader := repository.NewHotelReaderRepository(*db)
	hotelRepoWriter := repository.NewHotelWriterRepository(*db)

	// HotelEmp Repository
	hotelEmpRepoReader := emp.NewHotelEmpReaderRepository(*db)

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
		usecase.NewGetDetailHotelUseCase(hotelRepoReader),
		usecase.NewGetMyHotelUseCase(hotelRepoReader),
		usecase.NewUpdateHotelUseCase(hotelEmpRepoReader, hotelRepoWriter),
		usecase.NewGetNearMeHotelUseCase(hotelRepoReader),
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
			{
				Path:    "/filter/near-me",
				Method:  method.GET,
				Handler: r.HotelHandler.HandleGetNearMeHotel,
			},
			{
				Path:    "/:hotel_id",
				Method:  method.GET,
				Handler: r.HotelHandler.HandleGetDetailHotel,
			},
			{
				Path:    "/my-hotel",
				Method:  method.GET,
				Handler: r.HotelHandler.HandleGetMyHotel,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/update",
				Method:  method.PATCH,
				Handler: r.HotelHandler.HandleUpdateHotel,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
		},
	}
}
