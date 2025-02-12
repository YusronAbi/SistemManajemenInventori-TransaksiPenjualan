package handler

import (
	"encoding/json"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Authentication successful")
}
