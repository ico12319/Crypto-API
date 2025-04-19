package holding

import (
	"crptoApi/internal/utills"
	"crptoApi/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type HoldingService interface {
	GetHoldingRecord(id string) (models.Holding, error)
	GetHoldingsRecords() ([]models.Holding, error)
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
		utills.EncodeError(w, "error in the query parameter")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	holding, err := h.service.GetHoldingRecord(cryptoId)
	if err != nil {
		utills.EncodeError(w, "error when trying to get holding record")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(holding); err != nil {
		utills.EncodeError(w, "error when trying to encode holding JSON response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *HoldingHandler) GetHoldingsHandler(w http.ResponseWriter, r *http.Request) {
	holdings, err := h.service.GetHoldingsRecords()
	if err != nil {
		utills.EncodeError(w, "error when trying to get holdings records")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(holdings); err != nil {
		utills.EncodeError(w, "error when encoding holdings JSON response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
