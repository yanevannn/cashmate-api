package controllers

import (
	"cashmate-api/models"
	"cashmate-api/services"
	"cashmate-api/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var transaction models.CreateTransactionInput
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		utils.ResError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := transaction.Validate(); err != nil {
		utils.ResValidationError(w, err)
		return
	}

	userID := 1
	err := services.CreateTransactionsService(userID, transaction)
	if err != nil {
		utils.ResError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	utils.ResSuccess(w, http.StatusCreated, "Transaction Created Succesfuly", nil)
}

func GetTransactionByIdHandler (w http.ResponseWriter, r *http.Request) {
	// get transaction id from url
	transactionIdString := chi.URLParam(r, "id")
	transactionID, err := strconv.Atoi(transactionIdString)
	if err != nil || transactionID <= 0 {
		utils.ResError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	// get user id from token
	userID := 1 // for test purpose

	transaction, err := services.GetTransactionByIdService(transactionID, userID)
	if err != nil {
		utils.ResError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.ResSuccess(w, http.StatusOK, "Transaction retrieved successfully", transaction)
}

func UpdateTransactionHandler (w http.ResponseWriter, r *http.Request) {
	// get transaction id from url
	transactionIdString := chi.URLParam(r, "id")
	transactionID, err := strconv.Atoi(transactionIdString)
	if err != nil || transactionID <= 0 {
		utils.ResError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	// get user id from token
	userID := 1 // for test purpose

	var transaction models.UpdateTransactionInput
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		utils.ResError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := transaction.Validate(); err != nil {
		utils.ResValidationError(w, err)
		return
	}

	err = services.UpdateTransactionService(transactionID, userID, transaction)
	if err != nil {
		utils.ResError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	utils.ResSuccess(w, http.StatusOK, "Transaction updated successfully", nil)
}
