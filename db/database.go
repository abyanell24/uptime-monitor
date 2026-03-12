package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {

	connStr := "host=localhost port=5433 user=postgres password=postgres dbname=uptime sslmode=disable"

	database, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	DB = database

	fmt.Println("database connected")
}