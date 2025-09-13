package controllers

import (
	"cashmate-api/models"
	"cashmate-api/services"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
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

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Success: "false",
			Message: "Method not allowed",
		})
		return
	}

	// Extract user ID from query parameters
	idStr := strings.TrimPrefix(r.URL.Path, "/v1/user/")
    userID, err := strconv.Atoi(idStr)
    if err != nil || userID <= 0 {
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Success: "false",
			Message: "Invalid user ID" + err.Error(),
		})
		return
    }

	// Call Service to get user by ID
	user, err := services.GetUserByIDService(userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Success: "false",
			Message: err.Error(),
		})
		return
	}
	
	// Create a PublicUser instance to exclude the password
	publicUser := models.PublicUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: "true",
		Message: "User retrieved successfully",
		Data:    publicUser,
	})
}


func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Success: "false",
			Message: "Method not allowed",
		})
		return
	}

	users, err := services.GetAllUsersService()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Success: "false",
			Message: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: "true",
		Message: "Users retrieved successfully",
		Data:    users,
	})
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Success: "false",
			Message: "Method not allowed",
		})
		return
	}

	// Extract user ID from query parameters
	id := strings.TrimPrefix(r.URL.Path, "/v1/user/delete/")
	userID, err := strconv.Atoi(id)
	if err != nil || userID <= 0 {
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Success: "false",
			Message: "Invalid user ID",
		})
		return
	}

	// Call Service to delete user by ID
	err = services.DeleteUserService(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Success: "false",
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: "true",
		Message: "User deleted successfully",
	})
}
