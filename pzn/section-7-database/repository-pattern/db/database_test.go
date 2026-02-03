package db

import (
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestDBConn(t *testing.T) {
	db := GetConnection()

	db.Close()
}
