package helpers

import (
	"gin-training/serializers/responses"
)

type EmptyObj struct{}

func BuildDetailResponse(data interface{}) responses.DetailResponse {
	res := responses.DetailResponse{
		Status: true,
		Data:   data,
	}
	return res
}

func BuildListResponse(data interface{}, pagination interface{}) responses.ListResponse {
	res := responses.ListResponse{
		Status:     true,
		Data:       data,
		Pagination: pagination,
	}
	return res
}

func BuildErrorResponse(message string, err interface{}) responses.ErrorResponse {
	res := responses.ErrorResponse{
		Status:  false,
		Message: message,
		Errors:  err,
	}
	return res
}
