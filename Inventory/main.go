package main

import (
	"log"
	"net/http"
	"os"

	"inventory-system/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", handlers.GetProductByID).Methods("GET")
	r.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/transactions", handlers.GetTransactions).Methods("GET")
	r.HandleFunc("/transactions", handlers.CreateTransaction).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
