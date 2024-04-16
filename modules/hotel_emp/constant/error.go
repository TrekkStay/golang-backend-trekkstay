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

func ErrEmpNotFound(err error) error {
	return res.NewErrorResponse(
		http.StatusNotFound,
		err,
		"user employee not found",
		"ERR_EMPLOYEE_NOT_FOUND",
	)
}

func ErrPermissionDenied(err error) error {
	return res.NewErrorResponse(
		http.StatusNotFound,
		err,
		"permission denied",
		"ERR_PERMISSION_DENIED",
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

func ErrorPhoneAlreadyExists(err error) error {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"phone already exists",
		"ERR_PHONE_ALREADY_EXISTS",
	)
}

func ErrorWrongPassword(err error) error {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"wrong password",
		"ERR_WRONG_PASSWORD",
	)
}

func ErrNoHotelToAssignEmp(err error) error {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"no hotel to assign employee",
		"EER_NO_HOTEL_TO_ASSIGN_EMP",
	)
}

func ErrorHashPassword(err error) error {
	return res.NewErrorResponse(
		http.StatusInternalServerError,
		err,
		"hash password error",
		"ERR_HASH_PASSWORD",
	)
}
