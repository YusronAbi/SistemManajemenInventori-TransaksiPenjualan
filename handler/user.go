package handler

import (
	"encoding/json"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	users := []string{"Alice", "Bob", "Charlie"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
