package constant

import (
	"net/http"
	res "trekkstay/core/response"
)

func ErrCanNotCreateHotelRoom(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"can not create hotel room",
		"ERR_CAN_NOT_CREATE_HOTEL_ROOM",
	)
}
