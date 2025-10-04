package services

import (
	"cashmate-api/models"
	"cashmate-api/repositories"
	"errors"
	"time"
)

func GetAllTransactionsService(userID int) ([]models.Transaction, error) {
	transactions, err := repositories.GetAllTransactions(userID)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func CreateTransactionsService(userID int, Transaction models.CreateTransactionInput) error {
	// check category id type is exspense or income
	category, err := repositories.GetCategoryByID(Transaction.CategoriID)
	if err != nil {
		return err
	}
	transactionType := category.Type

	// Convert Date
	transactionDate, err := time.Parse("2006-01-02", Transaction.TransactionDate)
	if err != nil {
		return errors.New("Invalid date format. Please use YYYY-MM-DD")
	}
	Transaction.TransactionDate = transactionDate.Format("2006-01-02")

	// Create Transaction
	err = repositories.CreateTransaction(userID, transactionType, Transaction)
	if err != nil {
		return err
	}

	return nil
}

func GetTransactionByIdService(transactionID int, userID int) (*models.Transaction, error) {
	transaction, err := repositories.GetTransactionByID(transactionID, userID)
	if err != nil {
		return nil, errors.New("Transaction not found")
	}

	return &transaction, nil
}

func UpdateTransactionService(transactionID int, userID int, transaction models.UpdateTransactionInput) error {
	// check if transactionexists
	transactionExists, err := repositories.GetTransactionByID(transactionID, userID)
	if err != nil {
		return err
	}

	// check Type Category
	category, err := repositories.GetCategoryByID(transaction.CategoriID)
	if err != nil {
		return err
	}

	// check if the transaction belongs to the user
	if transactionExists.UserID != userID {
		return errors.New("You are not authorized to update this transaction")
	}

	// Update Transaction
	err = repositories.UpdateTransactionByID(transactionID, transaction, category.Type)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTransactionService(transactionID int, userID int) error {
	// check if transaction exists
	_, err := repositories.GetTransactionByID(transactionID, userID)
	if err != nil {
		return err
	}
	// Delete Transaction
	err = repositories.DeleteTransactionByID(transactionID, userID)
	if err != nil {
		return err
	}

	return nil
}
