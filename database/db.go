package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectToDb() {
	connStr := os.Getenv("CONNECTION_STRING")
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("err connecting db...", err)
	}

	defer db.Close()

	fmt.Println("Server running")

}
