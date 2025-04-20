package transaction

import (
	"context"
	"crptoApi/internal/utills"
	"crptoApi/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type TransactionService interface {
	CreateTransactionRecord(ctx context.Context, transaction models.Transaction) error
	GetTransactionRecord(id string) (models.Transaction, error)
	GetTransactionsRecords() ([]models.Transaction, error)
}

type TransactionHandler struct {
	service TransactionService
}

func NewTransactionHandler(service TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (t *TransactionHandler) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		utills.EncodeError(w, "invalid request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := t.service.CreateTransactionRecord(ctx, transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(transaction); err != nil {
		utills.EncodeError(w, "error when trying to encode transaction JSON response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (t *TransactionHandler) GetTransactionRecordHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		utills.EncodeError(w, "error in the query parameter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tModel, err := t.service.GetTransactionRecord(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(tModel); err != nil {
		utills.EncodeError(w, "error when trying to encode transaction JSON response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (t *TransactionHandler) GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	transactions, err := t.service.GetTransactionsRecords()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(&transactions); err != nil {
		utills.EncodeError(w, "error when trying to encode transaction JSON response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
