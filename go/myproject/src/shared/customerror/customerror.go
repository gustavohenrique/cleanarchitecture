package customerror

import (
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomError struct {
	code    codes.Code
	message string
}

const (
	UNKNOWN             = codes.Unknown            // 2
	INVALID             = codes.InvalidArgument    // 3
	NOT_FOUND           = codes.NotFound           // 5
	ALREADY_EXISTS      = codes.AlreadyExists      // 6
	PERMISSION_DENIED   = codes.PermissionDenied   // 7
	RESOURCE_EXHAUSTED  = codes.ResourceExhausted  // 8
	FAILED_PRECONDITION = codes.FailedPrecondition // 9
	INTERNAL            = codes.Internal           // 13
	UNAVAILABLE         = codes.Unavailable        // 14
	UNAUTHENTICATED     = codes.Unauthenticated    // 16
)

var httpCodes = map[codes.Code]int{
	codes.Unknown:            http.StatusConflict,
	codes.Internal:           http.StatusInternalServerError,
	codes.InvalidArgument:    http.StatusBadRequest,
	codes.NotFound:           http.StatusNotFound,
	codes.AlreadyExists:      http.StatusConflict,
	codes.Unauthenticated:    http.StatusUnauthorized,
	codes.Unavailable:        http.StatusInternalServerError,
	codes.PermissionDenied:   http.StatusForbidden,
	codes.FailedPrecondition: http.StatusPreconditionFailed,
	codes.ResourceExhausted:  http.StatusTooManyRequests,
}

var errorsMap = map[string]codes.Code{
	"violates":                    INVALID,
	"duplicate key":               ALREADY_EXISTS,
	"no rows in result":           NOT_FOUND,
	"no results found":            NOT_FOUND,
	"permission denied":           PERMISSION_DENIED,
	"no responders available":     UNAVAILABLE,
	"connect: connection refused": UNAVAILABLE,
	"connection":                  UNAVAILABLE,
	"empty":                       INVALID,
	"invalid token":               INVALID,
	"parse":                       INVALID,
	"mismatch":                    FAILED_PRECONDITION,
	"expired":                     PERMISSION_DENIED,
}

func (e *CustomError) Error() string {
	return e.message
}

func New(c codes.Code, m string) *CustomError {
	return &CustomError{
		code:    c,
		message: m,
	}
}

func Unavailable(message string) error {
	return Wrap(New(UNAVAILABLE, message))
}

func Denied(message string) error {
	return Wrap(New(PERMISSION_DENIED, message))
}

func PreCondition(message string) error {
	return Wrap(New(FAILED_PRECONDITION, message))
}

func Invalid(message string) error {
	return Wrap(New(INVALID, message))
}

func Unauthenticated(message string) error {
	return Wrap(New(UNAUTHENTICATED, message))
}

func NotFound(message string) error {
	return Wrap(New(NOT_FOUND, message))
}

func Wrap(err error) error {
	if err == nil {
		return nil
	}
	code := detect(err)
	return status.New(code, err.Error()).Err()
}

func detect(err error) codes.Code {
	ce, ok := err.(*CustomError)
	if !ok {
		return UNKNOWN
	}
	if ce.code > 0 {
		return ce.code
	}
	e := err.Error()
	for message, code := range errorsMap {
		if strings.Contains(strings.ToLower(e), strings.ToLower(message)) {
			return code
		}
	}
	return status.Code(err)
}

func StatusCodeFrom(err error) int {
	if err != nil {
		grpcCode, _ := status.FromError(err)
		httpCode := httpCodes[grpcCode.Code()]
		return httpCode
	}
	return http.StatusConflict
}
