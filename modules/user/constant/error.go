package constant

import (
	"net/http"
	res "trekkstay/core/response"
)

func ErrorEmailAlreadyExists(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"email already exists",
		"ERR_EMAIL_ALREADY_EXISTS",
	)
}

func ErrorInternalServerError(err error) error {
	return res.NewErrorResponse(
		http.StatusInternalServerError,
		err,
		"internal server error",
		"ERR_INTERNAL_SERVER_ERROR",
	)
}
