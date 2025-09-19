package services

import (
	"cashmate-api/models"
	"cashmate-api/repositories"
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
