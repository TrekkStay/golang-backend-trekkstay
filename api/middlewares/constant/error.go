package constant

import (
	"net/http"
	res "trekkstay/core/response"
)

func ErrInvalidToken(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusForbidden,
		err,
		"token is invalid signature",
		"ERR_INVALID_TOKEN",
	)
}

func ErrMissingToken(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusUnauthorized,
		err,
		"missing token in header",
		"ERR_UNAUTHORIZED",
	)
}

func ErrInternal(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusInternalServerError,
		err,
		"something went wrong with the server",
		"ERR_INTERNAL",
	)
}
