package account

import (
	"crptoApi/internal/utills"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//go:generate mockery --name=AccountService --output=./mocks --outpkg=mocks --filename=account_service.go --with-expecter=true
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encodedSuccessString := strconv.FormatFloat(balance, 'f', 2, 64) + "$"
	if err = json.NewEncoder(w).Encode(map[string]string{"balance": encodedSuccessString}); err != nil {
		utills.EncodeError(w, "error when trying to encode JSON response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *AccountHandler) UpdateBalanceHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	desiredBalance, ok := params["quantity"]
	if !ok {
		utills.EncodeError(w, "error with query parameter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	parsedBalance, err := strconv.ParseFloat(desiredBalance, 64)
	if err != nil {
		utills.EncodeError(w, "error when trying to parse provided quantity")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = a.service.UpdateAccountBalance(parsedBalance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	successString := buildSuccessStringHelper(w, parsedBalance)
	if err = json.NewEncoder(w).Encode(map[string]string{"success": successString}); err != nil {
		utills.EncodeError(w, "error when trying to encode JSON response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func buildSuccessStringHelper(w http.ResponseWriter, balance float64) string {
	var successString string
	if balance < 0 {
		successString = "Withdrew " + strconv.FormatFloat(balance, 'f', 2, 64) + "$"
	} else {
		successString = "Deposited " + strconv.FormatFloat(balance, 'f', 2, 64) + "$"
	}
	return successString
}
