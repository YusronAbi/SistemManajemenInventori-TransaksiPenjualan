package models

type Transaction struct {
	ID       int
	Product  Product
	Quantity int
	Total    float64
}
