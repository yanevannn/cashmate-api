package services

import (
	"cashmate-api/models"
	"cashmate-api/repositories"
	"fmt"
)

func GetAllCategoriesService(userID int) ([]models.Category, error) {
	categories, err := repositories.GetAllCategories(userID)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func CreateCategoryService(userID int, category *models.CreateCategoryInput) error {
	return repositories.CreateCategory(userID, category)
}

func GetCategoryByIDSevice(categoryID int) (*models.Category, error){
	category, err := repositories.GetCategoryByID(categoryID)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func UpdateCategoryService(category *models.UpdateCategoryInput, categoryID int, userID int) error {
	result, err := repositories.UpdateCategory(category, categoryID, userID)
	if err != nil {
		return err
	}

	// Check if any row was affected
	if result == 0 {
		return fmt.Errorf("category not found or not updated")
	}

	return nil
}

func DeleteCategoryService(categoryID int, userID int) error {
	category, err := repositories.GetCategoryByID(categoryID)
	if err != nil {
		return err
	}

	if userID != category.UserID {
		return fmt.Errorf("You are not authorized to delete this category")
	}

	err = repositories.DeleteCategory(categoryID, userID)
	if err != nil {
		return err
	}
	return nil
}
