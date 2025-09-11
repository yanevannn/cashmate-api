package models

type ErrorResponse struct {
	Success string `json:"success"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Success string      `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}