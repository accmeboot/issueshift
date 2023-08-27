package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("Failed to connect to db: %s", err.Error())
		return nil, err
	}

	return db, nil
}
