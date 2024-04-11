package constant

import (
	"net/http"
	res "trekkstay/core/response"
)

func ErrCanNotListDestination(err error) error {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"can not list destination",
		"ERR_CAN_NOT_LIST_DESTINATION",
	)
}

func ErrCanNotCreateDestination(err error) error {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"can not create destination",
		"ERR_CAN_NOT_CREATE_DESTINATION",
	)
}
