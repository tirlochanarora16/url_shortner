package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

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

func CreateTable(query, tableName string) {
	_, err := DB.Query(query)

	if err != nil {
		fmt.Println("error creating urls table", err)
		return
	}

	fmt.Println("Created table --- ", tableName)

}

func AlterTable(query, msg string) error {
	_, err := DB.Query(query)

	if err != nil {
		fmt.Println("error altering the table", err)
		return err
	}

	fmt.Println("Sucess Message --- ", msg)
	return nil
}

func RunMigrations() {
	for _, migration := range migrations {
		table, columnName, migrationName := migration.Table, migration.ColumnName, migration.MigrationName

		table = strings.TrimSpace(table)
		columnName = strings.TrimSpace(columnName)

		if table == "" || columnName == "" {
			fmt.Println("Table name or Column Name cannot be empty!")
			continue
		}

		_, err := checkColumnExistsInTable(table, columnName)

		if err != nil {
			fmt.Println("error in checking the column", err)
			continue
		}

		migrationApplied, err := migration.Check()

		if err != nil {
			fmt.Println("error checking the migration --- ", migrationName)
			continue
		}

		if migrationApplied {
			fmt.Println("Migration already applied...Skipping the migration ---", migrationName)
			continue
		} else {
			err := migration.Apply()

			if err != nil {
				fmt.Println(fmt.Errorf("Error in applying the migration '%s' on the Table '%s'", migrationName, table))
				continue
			}
		}
	}

	fmt.Println("Migrations complete... 😀")
}
