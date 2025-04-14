package holding

import (
	"crptoApi/pkg/constants"
	"crptoApi/pkg/models"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type HoldingService interface {
	GetHoldingRecord(id string) (models.Holding, error)
	GetHoldingsRecords() map[string]models.Holding
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
	holding, err := h.service.GetHoldingRecord(cryptoId)
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
	holdings := h.service.GetHoldingsRecords()
	if err := json.NewEncoder(w).Encode(holdings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
