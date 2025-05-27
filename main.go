package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tirlochanarora16/url_shortner/database"
	"github.com/tirlochanarora16/url_shortner/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error fetching env variables", err)
	}

	err = database.ConnectToDb()

	if err != nil {
		log.Fatal(err)
	}

	database.CreateUrlsTable()

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":3000")
}
