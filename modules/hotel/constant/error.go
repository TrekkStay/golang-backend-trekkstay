package constant

import (
	"net/http"
	res "trekkstay/core/response"
)

func ErrCanNotCreateHotel(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"can not create hotel",
		"ERR_CAN_NOT_CREATE_HOTEL",
	)
}

func ErrCantNotGetHotel(err error) error {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"cant not get hotel",
		"ERR_CANT_NOT_GET_HOTEL",
	)
}
