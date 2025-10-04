package repositories

import (
	"context"

	"cashmate-api/config"
	"cashmate-api/models"
)

func GetAllCategories(userID int) ([]models.Category, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	query := `SELECT id, user_id, name, type, description, icon, color, is_default, is_active, created_at, updated_at 
			  FROM categories 
			  WHERE is_default = TRUE OR user_id = $1 `
	rows, err := conn.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.UserID, &category.Name, &category.Type, &category.Description, &category.Icon, &category.Color, &category.IsDefault, &category.IsActive, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func CreateCategory(userID int, category *models.CreateCategoryInput) error {
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

func UpdateCategory(category *models.UpdateCategoryInput, categoryID int, userID int) (int64, error) {

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
			description = COALESCE($3, description), 
			icon = COALESCE($4, icon), 
			color = COALESCE($5, color), 
			updated_at = NOW()
		WHERE id = $6 AND user_id = $7
	`
	result, err := conn.Exec(context.Background(), query, category.Name, category.Type, category.Description, category.Icon, category.Color, categoryID, userID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

func GetCategoryByID(categoryID int) (*models.Category, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	var Category models.Category

	query := `SELECT id, user_id, name, type, description, icon, color, is_default, is_active, created_at, updated_at 
			  FROM categories WHERE id = $1`
	// If i wanna find 1 data can use QueryRow
	err = conn.QueryRow(context.Background(), query, categoryID).Scan(
		&Category.ID,
		&Category.UserID,
		&Category.Name,
		&Category.Type,
		&Category.Description,
		&Category.Icon,
		&Category.Color,
		&Category.IsDefault,
		&Category.IsActive,
		&Category.CreatedAt,
		&Category.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &Category, nil
}

func DeleteCategory(categoryID int, userID int) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	query := `DELETE FROM categories WHERE id = $1 AND user_id = $2 AND is_default = FALSE`
	_, err = conn.Exec(context.Background(), query, categoryID, userID)
	if err != nil {
		return err
	}

	return nil
}
