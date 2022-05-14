package main

import (
	"context"
	"interface_project/api"
	"interface_project/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// at first clear os environment variables,
// then load project environment variables from file in this path `utils/.env` into os env variables from
func init() {
	os.Clearenv()
	if err := godotenv.Load("utils/.env"); err != nil {
		log.Fatal("error while loading .env file.")
	}
}

func main() {
	client, ctx, cancel := database.GetContextAndClient() // create instances for client and ctx with cancel function from database package
	// calling cancel function of ctx(Context) at the end of main function scope.
	defer cancel()
	// calling close function of client instance at the end of main function scope.
	defer client.Close()
	// set ent client in debug mode.
	client = *client.Debug()
	// print schemas changes into terminal(stdout)
	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %+v", err)
	}
	// create ent schemas into database if there are not exists in db. this function uses context.Background.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema changes: %+v", err)
	}

	api.RunAPI(&ctx, &client)
}
