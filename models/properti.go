package models

import "time"

type Property struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Address     string    `json:"address"`
	Price       float64   `json:"price"`
	Status      string    `json:"status"`
	Area        float64   `json:"area"`
	Bedrooms    int       `json:"bedrooms"`
	Bathrooms   int       `json:"bathrooms"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
