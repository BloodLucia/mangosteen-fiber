package e

import "net/http"

func ErrInternalServer() *Error {
	return New(http.StatusInternalServerError, "ErrInternalServer")
}

func ErrBadRequest() *Error {
	return New(http.StatusBadRequest, "ErrBadRequest")
}

func ErrUnauthorized() *Error {
	return New(http.StatusUnauthorized, "ErrUnauthorized")
}

func ErrFormatJSON() *Error {
	return New(http.StatusUnprocessableEntity, "ErrFormatJSON")
}

func ErrForbidden() *Error {
	return New(http.StatusForbidden, "ErrForbidden")
}
