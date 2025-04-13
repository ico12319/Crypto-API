package controllers

import (
	"CryptoToken/internal/services/accountService"
	"CryptoToken/pkg/constants"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type AccountController struct {
	service accountService.AccountService
}

func NewAccountController(service accountService.AccountService) *AccountController {
	return &AccountController{service: service}
}

func (a *AccountController) GetBalance(w http.ResponseWriter, r *http.Request) {
	balance := a.service.GetAccountBalance()
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	if err := json.NewEncoder(w).Encode(balance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *AccountController) SetBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	desiredBalance, ok := params["quantity"]
	parsedBalance, err := strconv.ParseFloat(desiredBalance, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if !ok {
		http.Error(w, errors.New("invalid query parameter").Error(), http.StatusBadRequest)
	}
	if err := a.service.SetAccountBalance(parsedBalance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	if err := json.NewEncoder(w).Encode(parsedBalance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (a *AccountController) UpdateBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	balance, ok := params["quantity"]
	if !ok {
		http.Error(w, fmt.Errorf("invalid query parameter").Error(), http.StatusBadRequest)
	}
	pasedBalance, err := strconv.ParseFloat(balance, 64)
	if err != nil {
		http.Error(w, fmt.Errorf("invalid query parameter").Error(), http.StatusBadRequest)
	}
	if err := a.service.UpdateAccountBalance(pasedBalance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	if err := json.NewEncoder(w).Encode(pasedBalance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
