package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(dsn *string) {
	if dsn == nil {
		log.Fatal("DSN is nil")
	}

	var err error
	DB, err = sql.Open("postgres", *dsn)
	if err != nil {
		log.Fatalf("Unable to open DB: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v", err)
	}

	log.Println("✅ Successfully connected to PostgreSQL database")
}
