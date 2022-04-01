package models

type Response[T any] struct {
	ResponseCode int    `json:"response_code"`
	Message      string `json:"message"`
	Data         T      `json:"data"`
}
