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
