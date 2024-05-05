package routes

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"trekkstay/api/middlewares"
	"trekkstay/config"
	"trekkstay/config/models"
	hotelRoomHandler "trekkstay/modules/hotel_room/api/handler"
	"trekkstay/modules/hotel_room/domain/usecase"
	"trekkstay/modules/hotel_room/repository"
	database "trekkstay/pkgs/dbs/postgres"
	"trekkstay/pkgs/dbs/redis"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

func NewHotelRoomHandler(db *database.Database, requestValidator *validator.Validate) hotelRoomHandler.HotelRoomHandler {
	// Hotel Room Repository
	hotelRoomRepoReader := repository.NewHotelRoomReaderRepository(*db)
	hotelRoomRepoWriter := repository.NewHotelRoomWriterRepository(*db)

	// Redis
	redisConfig := config.LoadConfig(&models.RedisConfig{}).(*models.RedisConfig)
	var conn = redis.Connection{
		Address:  fmt.Sprint(redisConfig.RedisHost, ":", redisConfig.RedisPort),
		Password: redisConfig.RedisPassword,
		Database: redisConfig.RedisDB,
	}

	var redisInstance = redis.NewRedis(conn)

	return hotelRoomHandler.NewHotelRoomHandler(requestValidator, redisInstance,
		usecase.NewCreateHotelRoomUseCase(hotelRoomRepoWriter),
		usecase.NewFilterHotelRoomUseCase(hotelRoomRepoReader),
		usecase.NewUpdateHotelRoomUseCase(hotelRoomRepoWriter),
		usecase.NewGetDetailHotelRoomUseCase(hotelRoomRepoReader),
		usecase.NewDeleteHotelRoomUseCase(hotelRoomRepoWriter),
	)
}

func (r *RouteHandler) hotelRoomRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/hotel-room",
		Routes: []route.Route{
			{
				Path:    "/create",
				Method:  method.POST,
				Handler: r.HotelRoomHandler.HandleCreateHotelRoom,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/filter",
				Method:  method.GET,
				Handler: r.HotelRoomHandler.HandleFilterHotelRoom,
			},
			{
				Path:    "/update",
				Method:  method.PATCH,
				Handler: r.HotelRoomHandler.HandleUpdateHotelRoom,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/:hotel_room_id",
				Method:  method.GET,
				Handler: r.HotelRoomHandler.HandleGetDetailHotelRoom,
			},
			{
				Path:    "/:room_id",
				Method:  method.DELETE,
				Handler: r.HotelRoomHandler.HandleDeleteHotelRoom,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
		},
	}
}
