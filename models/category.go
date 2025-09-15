package models

type Category struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"` // "income" or "expense"
	Description *string `json:"description"`
	Icon        string  `json:"icon"`
	Color       string  `json:"color"`
	IsDefault   bool    `json:"is_default"`
	IsActive    bool    `json:"is_active"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type CreateCategoryInput struct {
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required,oneof=income expense"`
	Description string `json:"description,omitempty" validate:"omitempty"`
	Icon        string `json:"icon" validate:"required"`
	Color       string `json:"color" validate:"required"`
}

type UpdateCategoryInput struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name,omitempty" validate:"omitempty"`
	Type        string `json:"type,omitempty" validate:"omitempty,oneof=income expense"`
	Description string `json:"description,omitempty" validate:"omitempty"`
	Icon        string `json:"icon,omitempty" validate:"omitempty"`
	Color       string `json:"color,omitempty" validate:"omitempty"`
}
