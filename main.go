package main

import (
	"fmt"
	"inventory-sales/config"
	"inventory-sales/handlers"
	"inventory-sales/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Set up Gin router
	r := gin.Default()

	// Apply middleware
	r.Use(middleware.AuthMiddleware())

	// Routes
	r.GET("/products", handlers.GetProducts)
	r.POST("/products", handlers.AddProduct)
	r.POST("/transactions", handlers.AddTransaction)
	r.GET("/transactions", handlers.GetTransactions)

	// Start server with configured port
	port := config.AppConfig.ServerPort
	fmt.Printf("Server running on port %s\n", port)
	r.Run(":" + port)
}
