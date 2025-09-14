package repositories

import (
	"context"

	"cashmate-api/config"
	"cashmate-api/models"
)

func GetAllCategories() ([]models.Category, error) {
	userID := 1 // Temporary hardcoded user ID for testing purposes

	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	query := `SELECT id, name, type, description, icon, color, is_default, is_active, created_at::TEXT, updated_at::TEXT FROM categories WHERE user_id = $1`
	rows, err := conn.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Type, &category.Description, &category.Icon, &category.Color, &category.IsDefault, &category.IsActive, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func CreateCategory(category *models.CreateCategoryInput) error {
	userID := 1 // Temporary hardcoded user ID for testing purposes

	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	query := `INSERT INTO categories (name, type, description, icon, color, is_default, is_active, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = conn.Exec(context.Background(), query, category.Name, category.Type, category.Description, category.Icon, category.Color, false, true, userID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCategory(category *models.UpdateCategoryInput, userID int) (int64, error) {

	conn, err := config.ConnectDB()
	if err != nil {
		return 0, err
	}
	defer conn.Close(context.Background())

	query := `
		UPDATE categories 
		SET 
			name = COALESCE($1, name), 
			type = COALESCE($2, type),
			description = COALESCE($2, description), 
			icon = COALESCE($3, icon), 
			color = COALESCE($4, color), 
			updated_at = NOW()
		WHERE id = $5 AND user_id = $6
	`
	result, err := conn.Exec(context.Background(), query, category.Name, category.Type, category.Description, category.Icon, category.Color, category.ID, userID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
