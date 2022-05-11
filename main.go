package main

import (
	"context"
	"interface_project/api"
	"interface_project/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	os.Clearenv()
	if err := godotenv.Load("utils/.env"); err != nil {
		log.Fatal("error while loading .env file.")
	}
}

func main() {
	client, ctx, cancel := database.GetContextAndClient()
	defer cancel()
	defer client.Close()
	client = *client.Debug()

	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %+v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema changes: %+v", err)
	}

	api.RunAPI(&ctx, &client)

}
