package utills

import (
	"encoding/json"
	"net/http"
)

func EncodeError(w http.ResponseWriter, errorString string) {
	json.NewEncoder(w).Encode(map[string]string{"error": errorString})
}

func EncodeSuccess(w http.ResponseWriter, successString string) {
	json.NewEncoder(w).Encode(map[string]string{"success": successString})
}
