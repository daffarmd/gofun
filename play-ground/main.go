package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "host=10.8.1.15 port=5432 user=angondev password=S3cret!!! dbname=db_pdam_admin sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("DB connected successfully!")
	db.Close()
}
