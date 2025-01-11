package handlers

import (
	"net/http"

	"inventory-sales/models"
	"inventory-sales/services"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products := services.GetAllProducts()
	c.JSON(http.StatusOK, products)
}

func AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdProduct := services.AddProduct(product)
	c.JSON(http.StatusCreated, createdProduct)
}
