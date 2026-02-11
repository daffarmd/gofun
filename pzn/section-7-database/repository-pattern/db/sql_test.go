package db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

type Customer struct {
	ID        string
	Name      string
	Email     sql.NullString
	Balance   int
	Rating    float64
	CreatedAt time.Time
	BirthDate time.Time
	Married   bool
}

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

	sqlSelect := "SELECT id, name, email, balance, rating, created_at, birth_date, married FROM customer"
	rows, err := db.QueryContext(ctx, sqlSelect)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var result Customer

	for rows.Next() {
		err = rows.Scan(
			&result.ID,
			&result.Name,
			&result.Email,
			&result.Balance,
			&result.Rating,
			&result.CreatedAt,
			&result.BirthDate,
			&result.Married,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println("id     :", result.ID)
		fmt.Println("name   :", result.Name)
		if result.Email.Valid {
			fmt.Println("email  :", result.Email.String)
		}
		fmt.Println("balance:", result.Balance)
		fmt.Println("rating :", result.Rating)
		fmt.Println("married:", result.Married)
		fmt.Println("BirthDate:", result.BirthDate)
		fmt.Println("CreatedAt:", result.CreatedAt)
		fmt.Println("-----")
	}

	fmt.Println("Select Data Success")
}

func TestQuerySelectParam(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	usr := "admin"
	password := "admin"
	fmt.Println(usr)
	fmt.Println(password)

	sqlSelect := "SELECT username FROM usr_user WHERE username = $1 AND password = $2"
	fmt.Println(sqlSelect)
	rows, err := db.QueryContext(ctx, sqlSelect, usr, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hai %s Login kamu berhasil\n", username)
	} else {
		fmt.Println("Gagal Login")
	}

}

func TestLastId(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlInsert := `
		INSERT INTO comment (username, comment)
		VALUES ($1, $2)
		RETURNING id
	`

	var lastID int
	err := db.QueryRowContext(ctx, sqlInsert, "test", "test").Scan(&lastID)
	if err != nil {
		panic(err)
	}

	fmt.Println("Insert Data Success", lastID)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlInsert := `
		INSERT INTO comment (username, comment)
		VALUES ($1, $2)
		RETURNING id
	`

	statement, err := db.PrepareContext(ctx, sqlInsert)
	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "daffa" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke " + strconv.Itoa(i)

		var lastID string
		err := statement.QueryRowContext(ctx, email, comment).Scan(&lastID)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id ke ", lastID)
	}

}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	sqlInsert := `
		INSERT INTO comment (username, comment)
		VALUES ($1, $2)
		RETURNING id
	`

	for i := 0; i < 10; i++ {
		email := "daffa" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke " + strconv.Itoa(i)

		var lastID string
		err := tx.QueryRowContext(ctx, sqlInsert, email, comment).Scan(&lastID)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id ke ", lastID)
	}
	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}

}
