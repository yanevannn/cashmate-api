package controllers

import (
	"cashmate-api/models"
	"cashmate-api/services"
	"cashmate-api/utils"
	"encoding/json"
	"net/http"
)

func RegisterHandler (w http.ResponseWriter, r *http.Request) {
	var user models.RegisterUser
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ResError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validate input
	if err := user.Validate(); err != nil {
		utils.ResValidationError(w, err)
		return
	}

	// Call service to register user
	if err := services.RegisterUserService(&user); err != nil {
		utils.ResError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResSuccess(w, http.StatusCreated, "User registered successfully. Please check your email for the OTP to verify your account.", nil)

}

func LoginHandler (w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		utils.ResError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	// Validate input
	if err := loginRequest.Validate(); err != nil {
		utils.ResValidationError(w, err)
		return
	}

	// Call service to login user
	loginResponse, err := services.LoginUserService(&loginRequest)
	if err != nil {
		utils.ResError(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.ResSuccess(w, http.StatusOK, "Login successful", loginResponse)
}

// func RefreshTokenHandler (w http.ResponseWriter, r *http.Request) {

// }