package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	USER   = "postgres"
	PASS   = "secret"
	DBNAME = "postgres"
)

func Connect() *sql.DB {
	URL := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", USER, PASS, DBNAME)
	db, err := sql.Open("postgres", URL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func TestConnection() {
	con := Connect()
	defer con.Close()
	err := con.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")
}
