package services

import (
	"errors"
	"inventory-sales/models"
)

var transactions []models.Transaction
var transactionIDCounter = 1

func AddTransaction(productID, quantity int) (models.Transaction, error) {
	product, err := GetProductByID(productID)
	if err != nil {
		return models.Transaction{}, err
	}

	if product.Stock < quantity {
		return models.Transaction{}, errors.New("not enough stock")
	}

	total := float64(quantity) * product.Price
	transaction := models.Transaction{
		ID:       transactionIDCounter,
		Product:  product,
		Quantity: quantity,
		Total:    total,
	}
	transactionIDCounter++
	transactions = append(transactions, transaction)

	err = UpdateStock(productID, quantity)
	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}

func GetAllTransactions() []models.Transaction {
	return transactions
}
