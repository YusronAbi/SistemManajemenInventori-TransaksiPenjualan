package handler

import (
	"encoding/json"
	"net/http"
)

func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Transaction successful")
}
