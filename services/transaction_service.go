package services

import (
	"cashmate-api/models"
	"cashmate-api/repositories"
	"errors"
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

	// Create Transaction
	err = repositories.CreateTransasction(userID, transactionType, Transaction)
	if err != nil {
		return err
	}

	return nil
}

func GetTransactionByIdService (transactionID int, userID int) (*models.Transaction, error) {
	transaction, err := repositories.GetTransactionByID(transactionID, userID)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func UpdateTransactionService (transactionID int, userID int, transaction models.UpdateTransactionInput) error {
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