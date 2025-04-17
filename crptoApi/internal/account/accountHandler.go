package account

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type AccountService interface {
	GetAccountBalance() (float64, error)
	UpdateAccountBalance(amount float64) error
}
type AccountHandler struct {
	service AccountService
}

func NewAccountHandler(service AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

func (a *AccountHandler) GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	balance, err := a.service.GetAccountBalance()
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}
	if err := json.NewEncoder(w).Encode(balance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *AccountHandler) UpdateBalanceHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	desiredBalance, ok := params["quantity"]
	if !ok {
		http.Error(w, errors.New("invalid query parameter").Error(), http.StatusBadRequest)
		return
	}
	parsedBalance, err := strconv.ParseFloat(desiredBalance, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = a.service.UpdateAccountBalance(parsedBalance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
