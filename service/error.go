package service

import (
	"encoding/json"

	"github.com/best-expendables-v2/common-utils/util/validation"
)

type ValidationError []validation.ValidationError

func NewValidationError(err error) ValidationError {
	return validation.ParseValidationErr(err)
}

func (m ValidationError) Error() string {
	b, _ := json.Marshal(m)
	return string(b)
}

type ServiceError struct {
	Code    string
	Message string
}

func (f ServiceError) Error() string {
	if f.Message != "" {
		return f.Message
	}
	return f.Code
}

type ForbiddenError ServiceError

func (f ForbiddenError) Error() string {
	return ServiceError(f).Error()
}

type Unauthorized ServiceError

func (f Unauthorized) Error() string {
	return ServiceError(f).Error()
}

type NotFoundError ServiceError

func (f NotFoundError) Error() string {
	return ServiceError(f).Error()
}

type BadRequestError ServiceError

func (f BadRequestError) Error() string {
	return ServiceError(f).Error()
}

type InternalServerError ServiceError

func (f InternalServerError) Error() string {
	return ServiceError(f).Error()
}
