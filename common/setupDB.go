package common

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
)

func SetupDB() *sql.DB {
	err := godotenv.Load()
	CheckError(err)

	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
    CheckError(err)

    return db
}