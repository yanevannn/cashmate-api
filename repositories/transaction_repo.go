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
	query := "SELECT id, user_id, category_id, type, amount, description, transaction_date::TEXT, created_at::TEXT, updated_at::TEXT from transactions WHERE user_id = $1 AND deleted_at IS NULL"
	rows, err := conn.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var Transactions []models.Transaction
	for rows.Next() {
		var Transaction models.Transaction
		err := rows.Scan(&Transaction.ID, &Transaction.UserID, &Transaction.CategoriID, &Transaction.Type, &Transaction.Amount, &Transaction.Description, &Transaction.TransactionDate, &Transaction.CreatedAt, &Transaction.UpdatedAt)
		if err != nil {
			return nil, err
		}
		Transactions = append(Transactions, Transaction)
	}

	return Transactions, nil
}

func CreateTransasction(userID int, typeTransaction string, transaction models.CreateTransactionInput) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}

	defer conn.Close(context.Background())
	query := `INSERT INTO transactions 
			  (user_id, category_id, amount, type, description, transaction_date)
			  VALUES ($1, $2, $3, $4,$5, $6 )`
	_, err = conn.Exec(context.Background(), query, userID, transaction.CategoriID, transaction.Amount, typeTransaction, transaction.Description, transaction.TransactionDate)

	if err != nil {
		return err
	}
	return nil
}

func GetTransactionByID (transactionID int, userID int) (models.Transaction, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return models.Transaction{}, err
	}
	defer conn.Close(context.Background())
	query := "SELECT id, user_id, category_id, type, amount, description, transaction_date::TEXT, created_at::TEXT, updated_at::TEXT from transactions WHERE id = $1 AND user_id = $2 AND deleted_at IS NULL"
	row := conn.QueryRow(context.Background(), query, transactionID, userID)
	var Transaction models.Transaction
	err = row.Scan(&Transaction.ID, &Transaction.UserID, &Transaction.CategoriID, &Transaction.Type, &Transaction.Amount, &Transaction.Description, &Transaction.TransactionDate, &Transaction.CreatedAt, &Transaction.UpdatedAt)
	if err != nil {
		return models.Transaction{}, err
	}
	return Transaction, nil
}

func UpdateTransactionByID (transactionID int, transaction models.UpdateTransactionInput, categoryType string) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}

	defer conn.Close(context.Background())
	query := `UPDATE transactions 
			  SET category_id = $1, amount = $2, description = $3, transaction_date = $4, type = $5, updated_at = NOW()
			  WHERE id = $6`
	_, err = conn.Exec(context.Background(), query, transaction.CategoriID, transaction.Amount, transaction.Description, transaction.TransactionDate, categoryType, transactionID)

	if err != nil {
		return err
	}
	return nil
}

func DeleteTransactionByID (transactionID int, userID int) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}

	defer conn.Close(context.Background())
	query := `UPDATE transactions 
			  SET deleted_at = NOW()
			  WHERE id = $1 AND user_id = $2`
	_, err = conn.Exec(context.Background(), query, transactionID, userID)
	if err != nil {
		return err
	}
	return nil
}
