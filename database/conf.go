package database

import "os"

/// GetDatabaseConf get database parameters from os env that sets in ./main.go init function.
func GetDatabaseConf() *dataBaseConf {
	return &dataBaseConf{
		Username:     os.Getenv("USERNAME"),
		Password:     os.Getenv("PASSWORD"),
		Host:         os.Getenv("HOST"),
		DatabaseName: os.Getenv("DB_NAME"),
	}
}

type dataBaseConf struct {
	Username     string
	Password     string
	DatabaseName string
	Host         string
}
