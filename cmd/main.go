package main

import (
	"log"

	"github.com/ynadtochii/ecom/cmd/api"
	"github.com/ynadtochii/ecom/db"
	"github.com/ynadtochii/ecom/db/migrations"

	"github.com/joho/godotenv"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
}

func main() {

	db.Connect()

	migrations.Migrate()

	// db.Seed()

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
	log.Println("server running")
}
