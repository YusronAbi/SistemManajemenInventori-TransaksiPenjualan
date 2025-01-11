package inventory

import (
	"net/http"
	"strconv"

	"inventory-sales/models"

	"github.com/gin-gonic/gin"
)

func AddInventoryHandler(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdProduct := AddInventory(product)
	c.JSON(http.StatusCreated, createdProduct)
}

func GetInventoriesHandler(c *gin.Context) {
	inventories := GetAllInventories()
	c.JSON(http.StatusOK, inventories)
}

func UpdateInventoryHandler(c *gin.Context) {
	productID, err := strconv.Atoi(c.Query("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	newStock, err := strconv.Atoi(c.Query("stock"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock"})
		return
	}

	err = UpdateInventoryStock(productID, newStock)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "inventory stock updated"})
}
