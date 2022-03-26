package database

import (
	"context"
	"fmt"
	"interface_project/ent"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func GetDatabaseClient() *ent.Client {
	dbConf := GetDatabaseConf()
	if client, err := ent.Open("postgres", fmt.Sprintf(
		"user=%s dbname=%s host=%s password=%s sslmode=disable",
		dbConf.Username,
		dbConf.DatabaseName,
		dbConf.Host,
		dbConf.Password,
	)); err != nil {
		log.Fatal("on GetDatabaseClient() error happened: ", err)
		return nil
	} else {
		return client
	}
}

func GetDatabaseContext() (*context.Context, *context.CancelFunc) {
	// client, err := GetDatabaseClient()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer client.Close()

	// log.Println("Database Connected.")

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour*24)
	return &ctx, &cancel

}

func GetContextAndClient() (ent.Client, context.Context, context.CancelFunc) {
	client := GetDatabaseClient()
	ctx, cancel := GetDatabaseContext()
	return *client, *ctx, *cancel
}
