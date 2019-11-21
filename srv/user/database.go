package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	dbUser = "postgres"
	dbPwd = "1"
	dbHost = "localhost"
	dbName = "walkie"
)

func CreateConnection() (*gorm.DB, error) {

	// Get database details from environment variables
	//host := os.Getenv("DB_HOST")
	//user := os.Getenv("DB_USER")
	//DBName := os.Getenv("DB_NAME")
	//password := os.Getenv("DB_PASSWORD")

	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=disable password=%s",
			dbHost, dbUser, dbName, dbPwd,
		),
	)
}
