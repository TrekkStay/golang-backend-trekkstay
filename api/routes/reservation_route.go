package routes

import (
	"github.com/go-playground/validator/v10"
	"trekkstay/api/middlewares"
	"trekkstay/config"
	"trekkstay/config/models"
	hotel "trekkstay/modules/hotel/repository"
	room "trekkstay/modules/hotel_room/repository"
	"trekkstay/modules/reservation/api/handler"
	"trekkstay/modules/reservation/domain/usecase"
	"trekkstay/modules/reservation/repository"
	database "trekkstay/pkgs/dbs/postgres"
	"trekkstay/pkgs/s3"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

func NewReservationHandler(db *database.Database, requestValidator *validator.Validate) handler.ReservationHandler {
	// Reservation Repository
	reservationRepoReader := repository.NewReservationReaderRepository(*db)
	reservationRepoWriter := repository.NewReservationWriterRepository(*db)

	// HotelRoom Repository
	hotelRoomRepoReader := room.NewHotelRoomReaderRepository(*db)

	// Hotel Repository
	hotelRepoReader := hotel.NewHotelReaderRepository(*db)

	// S3
	s3Config := config.LoadConfig(&models.S3Config{}).(*models.S3Config)

	return handler.NewReservationHandler(
		requestValidator,
		usecase.NewCreateReservationUseCase(hotelRoomRepoReader,
			hotelRepoReader, reservationRepoWriter, s3.NewS3Upload(s3Config)),
		usecase.NewFilterReservationUseCase(reservationRepoReader),
		usecase.NewGetDetailReservationUseCase(reservationRepoReader),
		usecase.NewCancelReservationUseCase(reservationRepoWriter),
	)
}

func (r *RouteHandler) reservationRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/reservation",
		Routes: []route.Route{
			{
				Path:    "/create",
				Method:  method.POST,
				Handler: r.ReservationHandler.HandleCreateReservation,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/filter",
				Method:  method.GET,
				Handler: r.ReservationHandler.HandleFilterReservation,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/:reservation_id",
				Method:  method.GET,
				Handler: r.ReservationHandler.HandleGetDetailReservation,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/cancel/:reservation_id",
				Method:  method.DELETE,
				Handler: r.ReservationHandler.HandleCancelReservation,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
		},
	}
}
