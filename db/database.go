package db

import (
    "database/sql"
    "fmt"
    "os"

    _ "github.com/lib/pq"
)

var DB *sql.DB

// ConnectDB connects to the PostgreSQL database
func ConnectDB() {
    // Ambil DATABASE_URL dari environment variable
    dbURL := os.Getenv("DATABASE_URL")

    if dbURL == "" {
        // Kalau env belum diset → panic supaya kita tahu
        panic("DATABASE_URL environment variable not set")
    }

    var err error
    DB, err = sql.Open("postgres", dbURL)
    if err != nil {
        panic(err)
    }

    // Tes koneksi
    if err = DB.Ping(); err != nil {
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
    }

    fmt.Println("Database connected")
}