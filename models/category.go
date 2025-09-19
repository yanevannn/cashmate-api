package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"` // "income" or "expense"
	Description *string   `json:"description"`
	Icon        string    `json:"icon"`
	Color       string    `json:"color"`
	IsDefault   bool      `json:"is_default"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateCategoryInput struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"` // "income" or "expense"
	Description *string `json:"description,omitempty"`
	Icon        string  `json:"icon"`
	Color       string  `json:"color"`
}

func (c CreateCategoryInput) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.Type, validation.Required, validation.In("income", "expense")),
		validation.Field(&c.Description, validation.When(c.Description != nil, validation.Length(1, 255))),
		validation.Field(&c.Icon, validation.Required),
		validation.Field(&c.Color, validation.Required),
	)

}

type UpdateCategoryInput struct {
	Name        string  `json:"name,omitempty"`
	Type        string  `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
	Icon        string  `json:"icon,omitempty"`
	Color       string  `json:"color,omitempty"`
}

func (c UpdateCategoryInput) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.Type, validation.Required, validation.In("income", "expense")),
		validation.Field(&c.Icon, validation.Required),
		validation.Field(&c.Color, validation.Required),
	)
}
