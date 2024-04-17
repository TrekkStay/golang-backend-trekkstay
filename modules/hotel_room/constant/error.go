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

func ErrCanNotFilterHotelRoom(err error) error {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"can not filter hotel room",
		"ERR_CAN_NOT_FILTER_HOTEL_ROOM",
	)
}

func ErrHotelIDIsRequired(err error) error {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"hotel's id is required",
		"ERR_HOTEL_ID_IS_REQUIRED",
	)
}

func ErrSomethingWentWrong(err error) error {
	return res.NewErrorResponse(
		http.StatusInternalServerError,
		err,
		"something went wrong",
		"ERR_SOMETHING_WENT_WRONG",
	)
}
