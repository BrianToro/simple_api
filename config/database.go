package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectDB() (*gorm.DB, error) {
	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		log.Panicln("DB_HOST is required")
	}
	user, ok := os.LookupEnv("DB_USERNAME")
	if !ok {
		log.Panicln("DB_USERNAME is required")
	}
	password, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		log.Panicln("DB_PASSWORD is required")
	}
	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		log.Panicln("DB_NAME is required")
	}
	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		log.Panicln("DB_PORT is required")
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, user, password, dbName, port)
	conn, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
