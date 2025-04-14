package transaction

import (
	"crptoApi/pkg/constants"
	"crptoApi/pkg/models"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type TransactionService interface {
	CreateTransactionRecord(transaction models.Transaction) error
	GetTransactionRecord(id string) (models.Transaction, error)
	GetTransactionsRecords() map[string]models.Transaction
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
	}
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	if err := t.service.CreateTransactionRecord(transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := json.NewEncoder(w).Encode(transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (t *TransactionHandler) GetTransactionRecordHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		http.Error(w, errors.New("invalid query parameter").Error(), http.StatusBadRequest)
	}
	tModel, err := t.service.GetTransactionRecord(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	if err := json.NewEncoder(w).Encode(tModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (t *TransactionHandler) GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	transactions := t.service.GetTransactionsRecords()
	if err := json.NewEncoder(w).Encode(&transactions); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
