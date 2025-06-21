package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tirlochanarora16/url_shortner/database"
	"github.com/tirlochanarora16/url_shortner/pkg"
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

	err = pkg.InitRedis()

	if err != nil {
		log.Println("Error connecting to the redis DB", err)
		return
	}

	database.CreateTable(database.CreateUrlTable, "users")
	database.CreateTable(database.CreateSchemaMigrationTable, "schema")
	if os.Getenv("APPLY_MIGRATION") == "true" {
		database.RunMigrations()
	}

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":3000")
}
