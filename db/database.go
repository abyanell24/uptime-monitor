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
        panic("DATABASE_URL environment variable not set")
    }

    var err error
    DB, err = sql.Open("postgres", dbURL)
    if err != nil {
        panic(err)
    }

    if err = DB.Ping(); err != nil {
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
    }

    fmt.Println("Database connected")
}