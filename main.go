package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tirlochanarora16/url_shortner/database"
)

func main() {
	initiateStartup()
}

func initiateStartup() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error fetching env variables", err)
	}

	err = database.ConnectToDb()

	if err != nil {
		log.Fatal(err)
	}

	database.CreateUrlsTable()
}
