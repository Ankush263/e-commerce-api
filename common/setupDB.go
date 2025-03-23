package common

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
)

func SetupDB() *sql.DB {
	err := godotenv.Load()
	CheckError("Dotenv Error: ", err)

	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
    CheckError("Connetion Error: ", err)

    return db
}