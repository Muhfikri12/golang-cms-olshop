package database

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	connStr := "user=postgres dbname=db_olshop_lumo sslmode=disable password=admin host=localhost"
	db, err := sql.Open("postgres", connStr)

	if err := db.Ping(); err != nil {
		return nil, errors.New("failed to connect database")
	}

	return db, err
}