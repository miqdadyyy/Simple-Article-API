package entity

import "github.com/miqdadyyy/jetdevs-assesment/internal/constant"

type HTTPJsonResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewHTTPJsonResponse(status int, Data interface{}) HTTPJsonResponse {
	return HTTPJsonResponse{
		Status:  status,
		Message: constant.HTTPStatusMessageMap[status],
		Data:    Data,
	}
}
