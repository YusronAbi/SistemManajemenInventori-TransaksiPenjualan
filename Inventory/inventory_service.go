package inventory

import (
	"errors"
	"inventory-sales/models"
)

var inventories []models.Product
var inventoryIDCounter = 1

func AddInventory(product models.Product) models.Product {
	product.ID = inventoryIDCounter
	inventoryIDCounter++
	inventories = append(inventories, product)
	return product
}

func GetAllInventories() []models.Product {
	return inventories
}

func UpdateInventoryStock(productID, newStock int) error {
	for i, inventory := range inventories {
		if inventory.ID == productID {
			inventories[i].Stock = newStock
			return nil
		}
	}
	return errors.New("inventory not found")
}
