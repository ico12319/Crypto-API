package controllers

import (
	"CryptoToken/internal/services/transactionService"
	"CryptoToken/pkg/constants"
	"CryptoToken/pkg/models"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type TransactionController struct {
	service transactionService.TransactionService
}

func NewTransactionController(service transactionService.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

func (t *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err := t.service.CreateTransactionRecord(transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	if err := json.NewEncoder(w).Encode(transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (t *TransactionController) GetTransactionRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		http.Error(w, errors.New("invalid query parameter").Error(), http.StatusBadRequest)
	}
	transactionInstance, err := t.service.GetTransactionRecord(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	if err := json.NewEncoder(w).Encode(transactionInstance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (t *TransactionController) GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions := t.service.GetTransactions()
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	if err := json.NewEncoder(w).Encode(transactions); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
