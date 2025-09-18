package repositories

import (
	"cashmate-api/config"
	"cashmate-api/models"
	"context"
)

func GetAllTransactions(userID int) ([]models.Transaction, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())
	query := "SELECT id, category_id, type, amount, description, transaction_date::TEXT, created_at::TEXT, updated_at::TEXT from transactions WHERE user_id = $1"
	rows, err := conn.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var Transactions []models.Transaction
	for rows.Next() {
		var Transaction models.Transaction
		err := rows.Scan(&Transaction.ID, &Transaction.CategoriID, &Transaction.Type, &Transaction.Amount, &Transaction.Description, &Transaction.TransactionDate, &Transaction.CreatedAt, &Transaction.UpdatedAt)
		if err != nil {
			return nil, err
		}
		Transactions = append(Transactions, Transaction)
	}

	return Transactions, nil
}
