package utils

import (
	"cashmate-api/models"
	"encoding/json"
	"net/http"
)

func ResSuccess(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ResError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.APIResponse{
		Success: false,
		Message: message,
		Data:   nil,
	})
}