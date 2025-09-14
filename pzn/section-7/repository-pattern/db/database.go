package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Conn() {
	db, err := sql.Open("postgres", "devaja:1q2w3e4r@tcp(127.0.0.1:5432)/gofun")
	if err != nil {
		panic(err)
	}

	db.Close()
}
