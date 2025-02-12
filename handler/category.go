package handler

import (
	"ECOMERCE-API/models"
	"encoding/json"
	"net/http"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}
