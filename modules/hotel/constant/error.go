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
