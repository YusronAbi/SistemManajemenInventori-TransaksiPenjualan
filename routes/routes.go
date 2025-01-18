package routes

import (
	"SISTEMMANAJEMENINVENTORYTRANSAKSI/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/api/properties", func(c *gin.Context) {
		controllers.GetProperties(c, db)
	})

	router.POST("/api/transactions", func(c *gin.Context) {
		controllers.CreateTransaction(c, db)
	})

	router.Use(middleware.AuthMiddleware())
}
