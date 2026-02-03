package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	dsn := "postgres://daf-devs:1q2w3e4r@127.0.0.1:5432/gofun?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
