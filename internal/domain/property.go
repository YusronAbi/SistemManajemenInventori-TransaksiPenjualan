package domain

import (
	"time"
)

type Property struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"` // Rumah, Apartemen, dll
	Address     string    `json:"address"`
	Price       float64   `json:"price"`
	Status      string    `json:"status"` // Available, Sold, Reserved
	Area        float64   `json:"area"`
	Bedrooms    int       `json:"bedrooms"`
	Bathrooms   int       `json:"bathrooms"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Transaction struct {
	ID         int64     `json:"id"`
	PropertyID int64     `json:"property_id"`
	UserID     int64     `json:"user_id"`
	Type       string    `json:"type"`   // Sale, Rent
	Status     string    `json:"status"` // Pending, Completed, Cancelled
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
