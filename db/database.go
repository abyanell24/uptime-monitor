package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		panic("DATABASE_URL not set! Set it in Railway Environment Variables")
	}

	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to open DB: %v", err))
	}

	err = DB.Ping()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to DB: %v", err))
	}

	fmt.Println("Database connected")
}