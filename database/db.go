package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectToDb() error {
	connStr := os.Getenv("CONNECTION_STRING")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("err connecting db...", err)
		return err
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Failed to open DB connection")
		return err
	}

	fmt.Println("successfully connected to DB")

	DB = db

	return nil
}

func CreateUrlsTable() {
	query := `
		CREATE TABLE IF NOT EXISTS urls (
			id SERIAL PRIMARY KEY,
			short_code VARCHAR(10) UNIQUE NOT NULL,
			original_url TEXT NOT NULL,
			created_at TIMESTAMP  DEFAULT NOW()
		)
	`

	_, err := DB.Query(query)

	if err != nil {
		fmt.Println("error creating urls table", err)
		return
	}

	fmt.Println("Created Users table")

}
