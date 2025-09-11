package controllers

import (
	"cashmate-api/models"
	"cashmate-api/services"
	"encoding/json"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    // Cek Method Request [POST]
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) // Set status code 405 w.WriteHeader(number ex 404 or using http._status)
		json.NewEncoder(w).Encode(models.ErrorResponse{
            Success: "false",
			Message: "Method not allowed",
		})
		return
	}

    // Parse Body Request for preparing data to be inserted
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    // Call Service to Create User and send variable user
	if err := services.CreateUserService(&user); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Success: "false",
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: "true",
		Message: "User created successfully",
		Data:    user,
	})
}
