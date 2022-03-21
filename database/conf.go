package database

import "os"

/// GetDatabaseConf get database parameters from os env that sets in ./main.go init function.
func getDatabaseConf() *dataBaseConf {
	return &dataBaseConf{
		username:     os.Getenv("USERNAME"),
		password:     os.Getenv("PASSWORD"),
		host:         os.Getenv("HOST"),
		databaseName: os.Getenv("DB_NAME"),
	}
}

type dataBaseConf struct {
	username     string
	password     string
	databaseName string
	host         string
}
