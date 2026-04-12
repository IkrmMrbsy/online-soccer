package error

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrSqlError            = errors.New("database server failed to execute query")
	ErrTooManyRequests     = errors.New("too many requests")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInvalidToken        = errors.New("invalid token")
	ErrForbidden           = errors.New("forbidden")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrSqlError,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
}
