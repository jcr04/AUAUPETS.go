// internal/infrastructure/database/db.go

package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() error {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=12345 dbname=Pets host=localhost port=5432 sslmode=disable")
	return err
}

func GetDB() *sql.DB {
	return db
}
