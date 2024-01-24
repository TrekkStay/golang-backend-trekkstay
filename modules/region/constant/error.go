package constant

import (
	"net/http"
	res "trekkstay/core/response"
)

func ErrCannotGetWards(err error) error {
	return res.NewErrorResponse(
		http.StatusInternalServerError,
		err,
		"cannot get wards",
		"ERR_CANNOT_GET_WARDS",
	)
}

func ErrCannotGetDistricts(err error) error {
	return res.NewErrorResponse(
		http.StatusInternalServerError,
		err,
		"cannot get districts",
		"ERR_CANNOT_GET_DISTRICTS",
	)
}

func ErrCannotGetProvinces(err error) error {
	return res.NewErrorResponse(
		http.StatusInternalServerError,
		err,
		"cannot get provinces",
		"ERR_CANNOT_GET_PROVINCES",
	)
}
