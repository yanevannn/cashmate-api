package models

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Transaction struct {
	ID              int     `json:"id"`
	UserID          int     `json:"user_id"`
	CategoriID      int     `json:"category_id"`
	Type            string  `json:"transactypetion_date"`
	Amount          float64 `json:"amaount"`
	Description     *string `json:"description"`
	TransactionDate string  `json:"transaction_date"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
	DeletedAt       string  `json:"deleted_at"`
}

type CreateTransactionInput struct {
	CategoriID      int     `json:"category_id"`
	Amount          float64 `json:"amount"`
	Description     *string `json:"description"`
	TransactionDate string  `json:"transaction_date"`
}

func (c CreateTransactionInput) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.CategoriID, validation.Required),
		validation.Field(&c.Amount, validation.Required, validation.Min(0.01)),
		validation.Field(&c.Description, validation.When(c.Description != nil, validation.Length(1, 255))),
		validation.Field(&c.TransactionDate, validation.Required, validation.Date("2006-01-02")),
	)
}
