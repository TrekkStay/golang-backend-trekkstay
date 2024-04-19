package constant

import (
	"net/http"
	res "trekkstay/core/response"
)

func ErrCanNotListAttraction(err error) error {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"can not list attraction",
		"ERR_CAN_NOT_LIST_ATTRACTION",
	)
}

func ErrCanNotCreateAttraction(err error) error {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"can not create attraction",
		"ERR_CAN_NOT_CREATE_ATTRACTION",
	)
}
