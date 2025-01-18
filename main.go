package main

import (
	"log"
	"os"
)

func main() {
	db := config.InitDB()
	defer config.CloseDB(db)

	router := gin.Default()
	routes.SetupRoutes(router, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
