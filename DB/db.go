package db

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error

	dsn := "root:password@tcp(127.0.0.1:3306)/inventory"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}

	// Mengecek koneksi dengan database
	if err = DB.Ping(); err != nil {
		log.Fatal("Database connection error:", err)
	}

	log.Println("Database connected!")
}
