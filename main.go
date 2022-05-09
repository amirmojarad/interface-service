package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	os.Clearenv()
	if err := godotenv.Load("utils/.env"); err != nil {
		log.Fatal("error while loading .env file.")
	}
}

func main() {
	// client, ctx, cancel := database.GetContextAndClient()
	// defer cancel()
	// defer client.Close()
	// client = *client.Debug()

	// if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
	// 	log.Fatalf("failed printing schema changes: %+v", err)
	// }
	// if err := client.Schema.Create(context.Background()); err != nil {
	// 	log.Fatalf("failed creating schema changes: %+v", err)
	// }

	// api.RunAPI(&ctx, &client)

	strs := strings.Split("00:00:57,223 --\u003e 00:01:00,350", " ")
	for _, item := range strs {
		hms := strings.Split(item, ",")[0]
		millsecond := strings.Split(item, ",")[1]
		fmt.Println(hms + " " + millsecond)
	}
	// start, _ := time.ParseDuration(time.RFC3339)
	// fmt.Println(start)

}
