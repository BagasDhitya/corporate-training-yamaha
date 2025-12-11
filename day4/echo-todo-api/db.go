package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, sslMode)

	var err error
	DB, err = sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal("Failed to open DB : ", err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal("Failed to connect DB : ", err)
	}

	log.Println("Postgres Connected!")
}
