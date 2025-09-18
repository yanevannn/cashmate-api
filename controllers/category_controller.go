package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cashmate-api/models"
	"cashmate-api/services"
	"cashmate-api/utils"

	"github.com/go-chi/chi/v5"
)

func GetAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := services.GetAllCategoriesService()
	if err != nil {
		utils.ResError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResSuccess(w, http.StatusOK, "Categories retrieved successfully", categories)
}

func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
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

	if err := services.CreateCategoryService(&categoriesInput); err != nil {
		utils.ResError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResSuccess(w, http.StatusCreated, "Category created successfully", nil)
}

func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
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

	// Temporary hardcoded user ID for testing purposes
	userID := 1
	if err := services.UpdateCategoryService(&updateCategoryInput, categoriesID, userID); err != nil {
		utils.ResError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResSuccess(w, http.StatusOK, "Category updated successfully", nil)
}
