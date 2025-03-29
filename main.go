package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	router "github.com/ankush263/e-commerce-api/routers"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	fmt.Println("E-commerce API")

	// Database connection
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		// Default connection string with SSL mode explicitly disabled
		// Update these values according to your database configuration
		// connStr = "host=localhost user=postgres password=Ankush%40postgres dbname=e-commerce-api sslmode=disable"
		connStr = "host=localhost user=postgres password=Ankush@postgres263 dbname=e-commerce-api sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	// Test database connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	fmt.Println("Connected to database successfully")

	// Initialize router
	router := router.Router()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}