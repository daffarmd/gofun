package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	entity "github.com/daffarmd/gofun/pzn/section-7-database/repository-pattern/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (r *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlInsert := `
		INSERT INTO comment (username, comment)
		VALUES ($1, $2)
		RETURNING id
	`

	var lastID int
	err := r.DB.QueryRowContext(ctx, sqlInsert, comment.Username, comment.Comment).Scan(&lastID)
	if err != nil {
		return comment, err
	}

	comment.Id = lastID

	return comment, nil
}

func (r *commentRepositoryImpl) FindById(ctx context.Context, id int) (entity.Comment, error) {
	sqlSelect := "SELECT id, username, comment FROM comment WHERE id = $1 LIMIT 1"
	comment := entity.Comment{}
	rows, err := r.DB.QueryContext(ctx, sqlSelect, id)
	if err != nil {
		return comment, err
	}

	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&comment.Id, &comment.Username, &comment.Comment); err != nil {
			return comment, err
		}
		return comment, nil
	} else {
		return comment, errors.New("id " + strconv.Itoa(id) + "Not Found")
	}

}

func (r *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	sqlSelect := "SELECT id, username, comment FROM comment"

	rows, err := r.DB.QueryContext(ctx, sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		if err := rows.Scan(&comment.Id, &comment.Username, &comment.Comment); err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (r *commentRepositoryImpl) Update(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlUpdate := "UPDATE comment SET username = $1, comment = $2 WHERE id = $3"
	result, err := r.DB.ExecContext(ctx, sqlUpdate, comment.Username, comment.Comment, comment.Id)
	if err != nil {
		return comment, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return comment, err
	}

	if rowsAffected == 0 {
		return comment, errors.New("id not found")
	}

	return comment, nil

}

func (r *commentRepositoryImpl) Delete(ctx context.Context, id int) error {
	sqlDelete := "UPDATE comment SET status = $1 WHERE id = $2"
	result, err := r.DB.ExecContext(ctx, sqlDelete, -2, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("id not found")
	}

	return nil

}
