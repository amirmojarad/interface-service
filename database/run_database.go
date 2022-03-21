package database

import (
	"fmt"
	"interface_project/ent"

	_ "github.com/lib/pq"
)

func GetDatabaseClient() (*ent.Client, error) {
	dbConf := getDatabaseConf()
	return ent.Open("postgres", fmt.Sprintf("user=%s dbname=%s host=%s password=%s sslmode=disable", dbConf.username, dbConf.databaseName, dbConf.host, dbConf.password))
}
