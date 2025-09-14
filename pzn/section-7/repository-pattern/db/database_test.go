package db_test

import (
	"database/sql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestDBConn(t *testing.T) {
	db, err := sql.Open("postgresql", "devaja:1q2w3e4r@tcp(127.0.0.1:5432)/gofun")
	if err != nil {
		panic(err)
	}

	db.Close()
}
