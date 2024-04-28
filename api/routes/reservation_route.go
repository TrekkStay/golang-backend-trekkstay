package routes

import (
	"github.com/go-playground/validator/v10"
	"trekkstay/api/middlewares"
	"trekkstay/config"
	"trekkstay/config/models"
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

	// S3
	s3Config := config.LoadConfig(&models.S3Config{}).(*models.S3Config)

	return handler.NewReservationHandler(
		requestValidator,
		usecase.NewCreateReservationUseCase(hotelRoomRepoReader, reservationRepoWriter, s3.NewS3Upload(s3Config)),
		usecase.NewFilterReservationUseCase(reservationRepoReader),
		usecase.NewGetDetailReservationUseCase(reservationRepoReader),
	)
}

func (r *RouteHandler) reservationRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/reservation",
		Routes: []route.Route{
			{
				Path:    "/create",
				Method:  method.POST,
				Handler: r.ReservationHandle.HandleCreateReservation,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/filter",
				Method:  method.GET,
				Handler: r.ReservationHandle.HandleFilterReservation,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/:reservation_id",
				Method:  method.GET,
				Handler: r.ReservationHandle.HandleGetDetailReservation,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
		},
	}
}
