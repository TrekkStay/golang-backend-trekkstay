package mapper

import (
	"trekkstay/modules/reservation/api/model/req"
	"trekkstay/modules/reservation/domain/entity"
)

func ConvertCreateReservationReqToEntity(req req.CreateReservationReq) entity.ReservationEntity {
	return entity.ReservationEntity{
		RoomID:       req.RoomID,
		Quantity:     req.Quantity,
		CheckInDate:  req.CheckInDate,
		CheckOutDate: req.CheckOutDate,
		GuestInfo: entity.GuestInfoJSON{
			FullName: req.Guest.FullName,
			Contact:  req.Guest.Contact,
			Adults:   req.Guest.Adults,
			Children: req.Guest.Children,
		},
	}
}
