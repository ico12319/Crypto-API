package controllers

import (
	"CryptoToken/internal/services/holdingService"
	"CryptoToken/pkg/constants"
	"CryptoToken/pkg/models"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type HoldingController struct {
	service holdingService.HoldingService
}

func NewHoldingController(service holdingService.HoldingService) *HoldingController {
	return &HoldingController{service: service}
}

func (h *HoldingController) CreateHolding(w http.ResponseWriter, r *http.Request) {
	var holding models.Holding
	if err := json.NewDecoder(r.Body).Decode(&holding); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := h.service.CreateHoldingRecord(holding); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	if err := json.NewEncoder(w).Encode(holding); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *HoldingController) DeleteHolding(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cryptoId, ok := params["id"]
	if !ok {
		http.Error(w, errors.New("invalid query parameter").Error(), http.StatusBadRequest)
	}

	if err := h.service.DeleteHoldingRecord(cryptoId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *HoldingController) GetHoldingRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cryptoId, ok := params["id"]
	if !ok {
		http.Error(w, errors.New("invalid query parameter").Error(), http.StatusBadRequest)
	}

	holdingInstance, err := h.service.GetHoldingRecord(cryptoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	if err := json.NewEncoder(w).Encode(holdingInstance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *HoldingController) GetHoldings(w http.ResponseWriter, r *http.Request) {
	holdings := h.service.GetHoldingsRecords()
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	if err := json.NewEncoder(w).Encode(holdings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
