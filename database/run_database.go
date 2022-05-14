package database

import (
	"context"
	"fmt"
	"interface_project/ent"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// get database information that saved in os env variables.
// then return ent client.
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

// create context instance with its cancel function.
func GetDatabaseContext() (*context.Context, *context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour*24)
	return &ctx, &cancel
}

// package and call database client and context from Context package and return their instances.
func GetContextAndClient() (ent.Client, context.Context, context.CancelFunc) {
	client := GetDatabaseClient()
	ctx, cancel := GetDatabaseContext()
	return *client, *ctx, *cancel
}
