package routes

import (
	"github.com/gin-gonic/gin"
	"inventory-sales-api/controllers"
)

func SetupRouter(r *gin.Engine) {
	// Produk API Routes
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products", controllers.GetProducts)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
}
s