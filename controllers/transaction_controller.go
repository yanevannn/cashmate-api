package controllers

import (
	"cashmate-api/services"
	"cashmate-api/utils"
	"net/http"
)

func GetALlTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var userID int = 1
	transactions, err := services.GetAllTransactionsService(userID)
	if err != nil {
		utils.ResError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResSuccess(w, http.StatusOK, "Transactions retrieved successfully", transactions)
}
