package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func checkColumnExistsInTable(table, columnName string) (bool, error) {
	var columnExists bool
	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM information_schema.columns 
			WHERE table_name=$1 AND column_name=$2
		)
	`
	err := DB.QueryRow(query, table, columnName).Scan(&columnExists)

	if err != nil {
		return false, err
	}

	return columnExists, nil
}

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
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			short_code VARCHAR(10) UNIQUE NOT NULL,
			original_url TEXT NOT NULL,
			created_at TIMESTAMP  DEFAULT NOW(),
			updated_at TIMESTAMP  DEFAULT NOW()
		)
	`

	_, err := DB.Query(query)

	if err != nil {
		fmt.Println("error creating urls table", err)
		return
	}

	fmt.Println("Created Users table")

}

func AlterTable(query, msg string) error {
	_, err := DB.Query(query)

	if err != nil {
		fmt.Println("error altering the table", err)
		return err
	}

	fmt.Println("Sucess Message - ", msg)
	return nil
}

func RunMigrations() {
	for _, item := range migrations {
		table, columnName := item["table"], item["columnName"]
		res, err := checkColumnExistsInTable(table, columnName)

		if err != nil {
			fmt.Println("error in checking the column", err)
			continue
		}

		if res {
			fmt.Println("Column already exists in the given table...Skipping the migration")
			continue
		} else {
			err = AlterTable(item["query"], fmt.Sprintf("Updated table %s for column %s", table, columnName))

			if err != nil {
				fmt.Println(fmt.Sprintf("Error updating table %s for column %s - %s", table, columnName, err.Error()))
				continue
			}
		}
	}

	fmt.Println("Migrations complete...")
}
