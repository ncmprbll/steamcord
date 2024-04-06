package http_errors

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strings"
)

var (
	BadRequest            = "Bad request"
	WrongCredentials      = "Wrong Credentials"
	NotFound              = "Not Found"
	Unauthorized          = "Unauthorized"
	Forbidden             = "Forbidden"
	PermissionDenied      = "Permission Denied"
	ExpiredCSRFError      = "Expired CSRF token"
	WrongCSRFToken        = "Wrong CSRF token"
	CSRFNotPresented      = "CSRF not presented"
	NotRequiredFields     = "No such required fields"
	BadQueryParams        = "Invalid query params"
	InternalServerError   = "Internal Server Error"
	RequestTimeoutError   = "Request Timeout"
	ExistsEmailError      = "User with given email already exists"
	InvalidJWTToken       = "Invalid JWT token"
	InvalidJWTClaims      = "Invalid JWT claims"
	NotAllowedImageHeader = "Not allowed image header"
	NoCookie              = "not found cookie header"
	Conflict              = "Conflict"
	ServerError           = "Internal Server Error"
)

type IErrorWrapper interface {
	Status() int
	Error() string
}

type ErrorWrapper struct {
	Status int
	Error  string
}

func newErrorWrapper(s int, e string) *ErrorWrapper {
	return &ErrorWrapper{s, e}
}

func parseSQLErrors(err error) *ErrorWrapper {
	if errors.Is(err, sql.ErrNoRows) {
		return newErrorWrapper(http.StatusUnauthorized, Unauthorized)
	}

	if strings.Contains(err.Error(), "23505") {
		return newErrorWrapper(http.StatusConflict, Conflict)
	}

	return newErrorWrapper(http.StatusBadRequest, BadRequest)
}

func ErrorResponse(err error) *ErrorWrapper {
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return newErrorWrapper(http.StatusRequestTimeout, RequestTimeoutError)
	case strings.Contains(strings.ToLower(err.Error()), "sql"):
		return parseSQLErrors(err)
	case strings.Contains(err.Error(), "Unmarshal"):
		return newErrorWrapper(http.StatusBadRequest, BadRequest)
	case strings.Contains(err.Error(), "UUID"):
		return newErrorWrapper(http.StatusBadRequest, BadRequest)
	case strings.Contains(strings.ToLower(err.Error()), "cookie"):
		return newErrorWrapper(http.StatusUnauthorized, Unauthorized)
	case strings.Contains(strings.ToLower(err.Error()), "token"):
		return newErrorWrapper(http.StatusUnauthorized, Unauthorized)
	case strings.Contains(strings.ToLower(err.Error()), "bcrypt"):
		return newErrorWrapper(http.StatusUnauthorized, Unauthorized)
	default:
		return newErrorWrapper(http.StatusInternalServerError, ServerError)
	}
}
