package response

import (
	"net/http"

	"github.com/best-expendables-v2/common-utils/service"
)

func ConvertServiceError(err error) ApiResponse {
	switch err.(type) {
	case service.ValidationError:
		return CreateValidationErrResponse(err.(service.ValidationError))
	case service.ForbiddenError:
		return ErrorResponse(service.ServiceError(err.(service.ForbiddenError)), http.StatusForbidden)
	case service.NotFoundError:
		return ErrorResponse(service.ServiceError(err.(service.NotFoundError)), http.StatusNotFound)
	case service.Unauthorized:
		return ErrorResponse(service.ServiceError(err.(service.Unauthorized)), http.StatusUnauthorized)
	case service.BadRequestError:
		return ErrorResponse(service.ServiceError(err.(service.BadRequestError)), http.StatusBadRequest)
	}
	return ErrorResponse(err, http.StatusInternalServerError)
}
