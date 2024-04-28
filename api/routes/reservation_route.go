package routes

import (
	"github.com/go-playground/validator/v10"
	"trekkstay/api/middlewares"
	room "trekkstay/modules/hotel_room/repository"
	"trekkstay/modules/reservation/api/handler"
	"trekkstay/modules/reservation/domain/usecase"
	"trekkstay/modules/reservation/repository"
	database "trekkstay/pkgs/dbs/postgres"
	"trekkstay/pkgs/transport/http/method"
	"trekkstay/pkgs/transport/http/route"
)

func NewReservationHandler(db *database.Database, requestValidator *validator.Validate) handler.ReservationHandler {
	// Reservation Repository
	//reservationRepoReader := repository.NewReservationReaderRepository(*db)
	reservationRepoWriter := repository.NewReservationWriterRepository(*db)

	// HotelRoom Repository
	hotelRoomRepoReader := room.NewHotelRoomReaderRepository(*db)

	return handler.NewReservationHandler(
		requestValidator,
		usecase.NewCreateReservationUseCase(hotelRoomRepoReader, reservationRepoWriter),
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
		},
	}
}
