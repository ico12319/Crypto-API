package holding

import (
	"context"
	"crptoApi/pkg/constants"
	"crptoApi/pkg/models"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type HoldingService interface {
	GetHoldingRecord(ctx context.Context, id string) (models.Holding, error)
	GetHoldingsRecords(ctx context.Context) (map[string]models.Holding, error)
}

type HoldingHandler struct {
	service HoldingService
}

func NewHoldingHandler(service HoldingService) *HoldingHandler {
	return &HoldingHandler{service: service}
}

func (h *HoldingHandler) GetHoldingHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cryptoId, ok := params["crypto_id"]
	if !ok {
		http.Error(w, errors.New("invalid query parameter").Error(), http.StatusBadRequest)
	}
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	holding, err := h.service.GetHoldingRecord(ctx, cryptoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	if err := json.NewEncoder(w).Encode(holding); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *HoldingHandler) GetHoldingsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	holdings, err := h.service.GetHoldingsRecords(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
	}
	if err := json.NewEncoder(w).Encode(holdings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
