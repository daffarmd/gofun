package db

import (
	"context"
	"fmt"
	"testing"
)

func TestInserExecCtx(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlInsert := "INSERT INTO customer (id, name) VALUES ('df', 'daffa')"
	_, err := db.ExecContext(ctx, sqlInsert)
	if err != nil {
		panic(err)
	}

	fmt.Println("Insert Data Success")
}
func TestUpdateExecCtx(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlUpdate := "UPDATE customer SET id = 'test', name = 'test' where id = 'cct'"
	_, err := db.ExecContext(ctx, sqlUpdate)
	if err != nil {
		panic(err)
	}

	fmt.Println("Update Data Success")
}
func TestQuerySelect(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlSelect := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, sqlSelect)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id: ", id)
		fmt.Println("name: ", name)
	}

	fmt.Println("Select Data Success")
}
