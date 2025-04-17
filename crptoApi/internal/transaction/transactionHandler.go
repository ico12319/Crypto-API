package transaction

import (
	"context"
	"crptoApi/pkg/models"
	"encoding/json"
	"errors"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (t *TransactionHandler) GetTransactionRecordHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		http.Error(w, errors.New("invalid query parameter").Error(), http.StatusBadRequest)
		return
	}
	tModel, err := t.service.GetTransactionRecord(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(tModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (t *TransactionHandler) GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	transactions, err := t.service.GetTransactionsRecords()
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}
	if err := json.NewEncoder(w).Encode(&transactions); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
