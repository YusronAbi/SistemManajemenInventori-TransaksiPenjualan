package handler

import (
	"encoding/json"
	"net/http"
)

func ShippingHandler(w http.ResponseWriter, r *http.Request) {
	var shipping models.Shipping
	if err := json.NewDecoder(r.Body).Decode(&shipping); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt, err := database.DB.Prepare("INSERT INTO shipping(transaction_id, address, status) VALUES(?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(shipping.TransactionID, shipping.Address, "Pending")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Shipping initiated")
}
