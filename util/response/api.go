package response

import (
	"errors"
	"net/http"

	"github.com/best-expendables-v2/common-utils/service"
	"github.com/best-expendables-v2/common-utils/util/validation"
	"gorm.io/gorm"
)

type ApiResponse struct {
	Code     int               `json:"-"`
	Data     interface{}       `json:"data,omitempty"`
	Metadata Metadata          `json:"metadata,omitempty"`
	Headers  map[string]string `json:"-"`
	Errors   interface{}       `json:"errors,omitempty"`
}

func (a ApiResponse) AddMetadata(m Metadata) ApiResponse {
	a.Metadata = m
	return a
}

type Metadata struct {
	Total int `json:"total,omitempty"`
}

type Error struct {
	Message    string                       `json:"message,omitempty"`
	Code       string                       `json:"code,omitempty"`
	StatusCode int                          `json:"status_code,omitempty"`
	Errors     []validation.ValidationError `json:"errors,omitempty"`
}

// Ok 200-OK with data in body
func Ok(data interface{}) ApiResponse {
	return ApiResponse{
		Code: http.StatusOK,
		Data: data,
	}
}

// Accepted 202-Accepted with data in body
func Accepted(data interface{}) ApiResponse {
	return ApiResponse{
		Code: http.StatusAccepted,
		Data: data,
	}
}

// Created 201-Created with data in body
func Created(data interface{}) ApiResponse {
	return ApiResponse{
		Code: http.StatusCreated,
		Data: data,
	}
}

// PartialContent 206-Partial Content with data in body
func PartialContent(data interface{}) ApiResponse {
	return ApiResponse{
		Code: http.StatusPartialContent,
		Data: data,
	}
}

// BadRequest 400-Bad Request
func BadRequest(err error) ApiResponse {
	return ErrorResponse(err, http.StatusBadRequest)
}

// 500 Internal server error
func InternalServerError(err error) ApiResponse {
	return ErrorResponse(err, http.StatusInternalServerError)
}

// ValidationError 422-Unprocessable Enttity
func ValidationError(err error) ApiResponse {
	errMessages := validation.ParseValidationErr(err)
	var response ApiResponse
	response.Code = http.StatusUnprocessableEntity
	response.Errors = Error{
		Message:    http.StatusText(http.StatusUnprocessableEntity),
		StatusCode: http.StatusUnprocessableEntity,
		Errors:     errMessages,
	}
	return response
}

// ErrorResponse Custom error response
func ErrorResponse(err error, status int) ApiResponse {
	var response ApiResponse
	response.Code = status
	errors := Error{
		Message:    err.Error(),
		StatusCode: status,
	}
	if errors.Message == "" {
		errors.Message = http.StatusText(status)
	}
	if err, ok := err.(service.ServiceError); ok {
		if err.Code != "" {
			errors.Code = err.Code
		} else {
			errors.Code = service.GetDefaultErrorCode(status)
		}
	}
	response.Errors = errors
	return response
}

func RecordNotFoundError(err error, message string) ApiResponse {
	if err == gorm.ErrRecordNotFound {
		return ErrorResponse(errors.New(message), http.StatusNotFound)
	}

	return ErrorResponse(err, http.StatusInternalServerError)
}

func CreateValidationErrResponse(errMessages []validation.ValidationError) ApiResponse {
	var response ApiResponse
	response.Code = http.StatusUnprocessableEntity
	response.Errors = Error{
		Message:    http.StatusText(http.StatusUnprocessableEntity),
		StatusCode: http.StatusUnprocessableEntity,
		Errors:     errMessages,
	}
	return response
}
