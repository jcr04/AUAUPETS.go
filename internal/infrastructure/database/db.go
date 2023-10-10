// internal/infrastructure/database/db.go

package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() error {
	var err error
	db, err = sql.Open("postgres", "user=postgres dbname=Pets sslmode=disable")
	return err
}

func GetDB() *sql.DB {
	return db
}
