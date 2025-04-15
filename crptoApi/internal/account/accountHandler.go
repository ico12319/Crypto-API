package account

import (
	"context"
	"crptoApi/pkg/constants"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type AccountService interface {
	GetAccountBalance(ctx context.Context) (float64, error)
	UpdateAccountBalance(ctx context.Context, amount float64) error
}
type AccountHandler struct {
	service AccountService
}

func NewAccountHandler(service AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

func (a *AccountHandler) GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	balance, err := a.service.GetAccountBalance(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
	}
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
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	if err = a.service.UpdateAccountBalance(ctx, parsedBalance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
