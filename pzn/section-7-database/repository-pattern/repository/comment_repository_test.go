package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/daffarmd/gofun/pzn/section-7-database/repository-pattern/db"
	entity "github.com/daffarmd/gofun/pzn/section-7-database/repository-pattern/entity"
)

func TestCommentInsert(t *testing.T) {
	dbConn := db.GetConnection()
	defer dbConn.Close()

	repo := NewCommentRepository(dbConn)
	ctx := context.Background()

	coment := entity.Comment{
		Username: "daffa",
		Comment:  "hello",
	}

	result, err := repo.Insert(ctx, coment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func TestCommentFindById(t *testing.T) {
	dbConn := db.GetConnection()
	defer dbConn.Close()

	repo := NewCommentRepository(dbConn)
	ctx := context.Background()

	comment, err := repo.FindById(ctx, 38)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)

}

func TestCommentFindAll(t *testing.T) {
	dbConn := db.GetConnection()
	defer dbConn.Close()

	repo := NewCommentRepository(dbConn)
	ctx := context.Background()

	comment, err := repo.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)

}

func TestUpdate(t *testing.T) {
	dbConn := db.GetConnection()
	defer dbConn.Close()

	repo := NewCommentRepository(dbConn)
	ctx := context.Background()
	coment := entity.Comment{
		Id:       38,
		Username: "daffa 1",
		Comment:  "hello test update",
	}

	comment, err := repo.Update(ctx, coment)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestDelete(t *testing.T) {
	dbConn := db.GetConnection()
	defer dbConn.Close()

	repo := NewCommentRepository(dbConn)
	ctx := context.Background()

	err := repo.Delete(ctx, 38)
	if err != nil {
		t.Fatalf("delete failed: %v", err)
	}

	fmt.Println("Data Berhasil di hapus")

}
