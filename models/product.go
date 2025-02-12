package models

import "time"

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CategoryID  int64     `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductRepository interface {
	GetByID(id int64) (*Product, error)
	Store(product *Product) error
	Update(product *Product) error
	Delete(id int64) error
	Fetch(cursor string, num int64) ([]Product, string, error)
}

type ProductUsecase interface {
	GetByID(id int64) (*Product, error)
	Store(product *Product) error
	Update(product *Product) error
	Delete(id int64) error
	Fetch(cursor string, num int64) ([]Product, string, error)
}
