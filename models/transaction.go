package models

import "time"

type Transaction struct {
	ID         int64     `json:"id"`
	PropertyID int64     `json:"property_id"`
	UserID     int64     `json:"user_id"`
	Type       string    `json:"type"`
	Status     string    `json:"status"`
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
