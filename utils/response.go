package utils

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
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

	errorsMap := make(map[string]string)

	// cek apakah error dari ozzo-validation
	if ve, ok := err.(validation.Errors); ok {
		for field, e := range ve {
			errorsMap[field] = e.Error()
		}
	} else {
		// fallback untuk error biasa
		errorsMap["error"] = err.Error()
	}

	json.NewEncoder(w).Encode(ValidationErrorResponse{
		Success: false,
		Message: "Validation failed",
		Errors:  errorsMap,
	})
}
