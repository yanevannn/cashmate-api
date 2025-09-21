package controllers

import (
	"net/http"
	"strconv"

	"cashmate-api/models"
	"cashmate-api/services"
	"cashmate-api/utils"

	"github.com/go-chi/chi/v5"
)

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from query parameters
	idStr := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(idStr)
	if err != nil || userID <= 0 {
		utils.ResError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Call Service to get user by ID
	user, err := services.GetUserByIDService(userID)
	if err != nil {
		utils.ResError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Create a PublicUser instance to exclude the password
	publicUser := models.PublicUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	utils.ResSuccess(w, http.StatusOK, "User retrieved successfully", publicUser)
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllUsersService()
	if err != nil {
		utils.ResError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResSuccess(w, http.StatusOK, "Users retrieved successfully", users)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from query parameters
	id := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(id)
	if err != nil || userID <= 0 {
		utils.ResError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Call Service to delete user by ID
	err = services.DeleteUserService(userID)
	if err != nil {
		utils.ResError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResSuccess(w, http.StatusOK, "User deleted successfully", nil)
}
