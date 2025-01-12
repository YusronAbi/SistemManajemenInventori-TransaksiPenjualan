package main

import (
	"fmt"
	"inventory-sales/config"
	"inventory-sales/handlers"
	"inventory-sales/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	r.GET("/products", handlers.GetProducts)
	r.POST("/products", handlers.AddProduct)
	r.POST("/transactions", handlers.AddTransaction)
	r.GET("/transactions", handlers.GetTransactions)

	port := config.AppConfig.ServerPort
	fmt.Printf("Server running on port %s\n", port)
	r.Run(":" + port)
}
