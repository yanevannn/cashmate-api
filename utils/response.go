package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ValidationErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func ResSuccess(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ResError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(APIResponse{
		Success: false,
		Message: message,
		Data:    nil,
	})
}

func ResValidationError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	if errs, ok := err.(validator.ValidationErrors); ok {
		errorsMap := make(map[string]string)
		for _, e := range errs {
			errorsMap[e.Field()] = e.Error()
		}

		json.NewEncoder(w).Encode(ValidationErrorResponse{
			Success: false,
			Message: "Validation failed",
			Errors:  errorsMap,
		})
		return
	}
}
