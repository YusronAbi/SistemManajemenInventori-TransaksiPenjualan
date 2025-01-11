package handlers

import (
	"net/http"
	"strconv"

	"inventory-sales/services"

	"github.com/gin-gonic/gin"
)

func AddTransaction(c *gin.Context) {
	productID, err := strconv.Atoi(c.Query("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	quantity, err := strconv.Atoi(c.Query("quantity"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
		return
	}

	transaction, err := services.AddTransaction(productID, quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

func GetTransactions(c *gin.Context) {
	transactions := services.GetAllTransactions()
	c.JSON(http.StatusOK, transactions)
}
