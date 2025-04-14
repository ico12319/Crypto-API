package account

import (
	"crptoApi/pkg/constants"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type AccountService interface {
	GetAccountBalance() float64
	UpdateAccountBalance(amoun float64) error
}
type AccountHandler struct {
	service AccountService
}

func NewAccountHandler(service AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

func (a *AccountHandler) GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	balance := a.service.GetAccountBalance()
	if err := json.NewEncoder(w).Encode(balance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *AccountHandler) UpdateBalanceHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	desiredBalance, ok := params["quantity"]
	if !ok {
		http.Error(w, errors.New("invalid query parameter").Error(), http.StatusBadRequest)
	}
	parsedBalance, err := strconv.ParseFloat(desiredBalance, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	if err = a.service.UpdateAccountBalance(parsedBalance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
