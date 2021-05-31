package service

import "net/http"

var defaultHttpErrorCode = map[int]string{
	http.StatusForbidden:           "Forbidden",
	http.StatusNotFound:            "NotFound",
	http.StatusUnauthorized:        "Unauthorized",
	http.StatusBadRequest:          "BadRequest",
	http.StatusInternalServerError: "InternalServerError",
}

func GetDefaultErrorCode(statusCode int) string {
	return defaultHttpErrorCode[statusCode]
}
