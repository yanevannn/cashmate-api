package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cashmate-api/middlewares"
	"cashmate-api/models"
	"cashmate-api/services"
	"cashmate-api/utils"

	"github.com/go-chi/chi/v5"
)

func GetAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	// Get user id from token
	claims, ok := middlewares.GetClaimsFromCtx(r)
	if !ok {
		utils.ResError(w, http.StatusUnauthorized, "Missing or invalid token claims")
		return
	}
	userID := claims.UserID

	categories, err := services.GetAllCategoriesService(userID)

	if err != nil {
		utils.ResError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResSuccess(w, http.StatusOK, "Categories retrieved successfully", categories)
}

func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	// Check user Id from token
	claims, ok := middlewares.GetClaimsFromCtx(r)
	if !ok {
		utils.ResError(w, http.StatusUnauthorized, "Missing or invalid token claims")
		return
	}
	userID := claims.UserID

	defer r.Body.Close()
	var categoriesInput models.CreateCategoryInput
	if err := json.NewDecoder(r.Body).Decode(&categoriesInput); err != nil {
		utils.ResError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validasi input menggunakan ozzo
	if err := categoriesInput.Validate(); err != nil {
		utils.ResValidationError(w, err) // Pastikan fungsi ini bisa menampilkan error ozzo
		return
	}

	if err := services.CreateCategoryService(userID, &categoriesInput); err != nil {
		utils.ResError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResSuccess(w, http.StatusCreated, "Category created successfully", nil)
}

func GetCategoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from query parameters
	idString := chi.URLParam(r, "id")
	categoryID, err := strconv.Atoi(idString)
	if err != nil || categoryID <= 0 {
		utils.ResError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	
	category, err := services.GetCategoryByIDSevice(categoryID)
	if err != nil {
		utils.ResError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	utils.ResSuccess(w, http.StatusOK, "Category retrieved successfully", category)
}

func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := middlewares.GetClaimsFromCtx(r)
	if !ok {
		utils.ResError(w, http.StatusUnauthorized, "Missing or invalid token claims")
		return
	}
	userID := claims.UserID

	// Extract user ID from query parameters
	idStr := chi.URLParam(r, "id")
	categoriesID, err := strconv.Atoi(idStr)
	if err != nil || categoriesID <= 0 {
		utils.ResError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	defer r.Body.Close()
	var updateCategoryInput models.UpdateCategoryInput
	if err := json.NewDecoder(r.Body).Decode(&updateCategoryInput); err != nil {
		utils.ResError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validasi input menggunakan ozzo
	if err := updateCategoryInput.Validate(); err != nil {
		utils.ResValidationError(w, err) // Pastikan fungsi ini bisa menampilkan error ozzo
		return
	}

	// Call the service to update the category
	if err := services.UpdateCategoryService(&updateCategoryInput, categoriesID, userID); err != nil {
		utils.ResError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResSuccess(w, http.StatusOK, "Category updated successfully", nil)
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := middlewares.GetClaimsFromCtx(r)
	if !ok {
		utils.ResError(w, http.StatusUnauthorized, "Missing or invalid token claims")
		return
	}
	userID := claims.UserID

	// Extract user ID from query parameters
	idStr := chi.URLParam(r, "id")
	categoriesID, err := strconv.Atoi(idStr)
	if err != nil || categoriesID <= 0 {
		utils.ResError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	// Call the service to delete the category
	if err := services.DeleteCategoryService(categoriesID, userID); err != nil {
		utils.ResError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResSuccess(w, http.StatusOK, "Category deleted successfully", nil)
}
