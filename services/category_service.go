package services

import (
	"cashmate-api/models"
	"cashmate-api/repositories"
	"fmt"
)


func GetAllCategoriesService() ([]models.Category, error){
	categories, err := repositories.GetAllCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func CreateCategoryService(category *models.CreateCategoryInput,) error {
	return repositories.CreateCategory(category)
}

func UpdateCategoryService(category *models.UpdateCategoryInput,id int, userID int) error {
	if id <= 0 {
		return fmt.Errorf("invalid category ID")
	}
	category.ID = id

	result, err := repositories.UpdateCategory(category, userID)
	if err != nil {
		return err
	}
	
	// Check if any row was affected
	if result == 0 {
		return fmt.Errorf("category not found or not updated")
	}

	return nil
}