package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CategoryID  int64   `json:"category_id"`
}

type ProductHandler struct {
	db *sql.DB
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{db}
}

func (h *ProductHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/products", h.CreateProduct).Methods("POST")
	r.HandleFunc("/products", h.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", h.GetProduct).Methods("GET")
	r.HandleFunc("/products/{id}", h.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", h.DeleteProduct).Methods("DELETE")
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var product Product
	err = h.db.QueryRow("SELECT id, name, description, price, stock, category_id FROM products WHERE id = ?", id).
		Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID)

	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, "Product not found")
		return
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, product)
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query("SELECT id, name, description, price, stock, category_id FROM products")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CategoryID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		products = append(products, p)
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	result, err := h.db.Exec("INSERT INTO products (name, description, price, stock, category_id) VALUES (?, ?, ?, ?, ?)",
		product.Name, product.Description, product.Price, product.Stock, product.CategoryID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	product.ID = id

	respondWithJSON(w, http.StatusCreated, product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	_, err = h.db.Exec("UPDATE products SET name=?, description=?, price=?, stock=?, category_id=? WHERE id=?",
		product.Name, product.Description, product.Price, product.Stock, product.CategoryID, id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	product.ID = id
	respondWithJSON(w, http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	result, err := h.db.Exec("DELETE FROM products WHERE id=?", id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected == 0 {
		respondWithError(w, http.StatusNotFound, "Product not found")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Product deleted successfully"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
