package main

import (
	"context"
	"interface_project/database"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	os.Clearenv()
	if err := godotenv.Load("utils/.env"); err != nil {
		log.Fatal("error while loading .env file.")
	}
}

func main() {
	client, err := database.GetDatabaseClient()
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()
	client = client.Debug()

	log.Println("Database Connected.")

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour*24)
	defer cancel()

	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %+v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema changes: %+v", err)
	}
}
