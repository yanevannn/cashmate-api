package models

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"` // "income" or "expense"
	Description *string `json:"description"`
	Icon string `json:"icon"`
	Color string `json:"color"`
	IsDefault bool `json:"is_default"`
	IsActive bool `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateCategoryInput struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"` // "income" or "expense"
	Description string `json:"description,omitempty"`
	Icon string `json:"icon"`
	Color string `json:"color"`
}

type UpdateCategoryInput struct {
	ID int `json:"id"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"` // "income" or "expense"
	Description string `json:"description,omitempty"`
	Icon string `json:"icon,omitempty"`
	Color string `json:"color,omitempty"`
}