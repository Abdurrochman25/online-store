package common

import "net/http"

const (
	OK = string('0'+(iota+1)/100%10) + string('0'+(iota+1)/10%10) + string('1'+(iota)/1%10)
	BadRequest
	Unauthorized
	Forbidden
	NotFound
	InternalServerError
)

var mapErrorCode = map[int]string{
	http.StatusOK:                  OK,
	http.StatusBadRequest:          BadRequest,
	http.StatusUnauthorized:        Unauthorized,
	http.StatusForbidden:           Forbidden,
	http.StatusNotFound:            NotFound,
	http.StatusInternalServerError: InternalServerError,
}

func GetErrorCode(statusCode int) string {
	return mapErrorCode[statusCode]
}
