package services

import (
	"errors"
	"inventory-sales/models"
)

var products []models.Product
var productIDCounter = 1

func AddProduct(product models.Product) models.Product {
	product.ID = productIDCounter
	productIDCounter++
	products = append(products, product)
	return product
}

func GetAllProducts() []models.Product {
	return products
}

func GetProductByID(id int) (models.Product, error) {
	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}
	return models.Product{}, errors.New("product not found")
}

func UpdateStock(productID, quantity int) error {
	for i, product := range products {
		if product.ID == productID {
			if product.Stock < quantity {
				return errors.New("stock is not sufficient")
			}
			products[i].Stock -= quantity
			return nil
		}
	}
	return errors.New("product not found")
}
