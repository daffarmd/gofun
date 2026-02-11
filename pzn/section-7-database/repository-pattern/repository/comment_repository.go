package repository

import (
	"context"

	entity "github.com/daffarmd/gofun/pzn/section-7-database/repository-pattern/entity"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
	Update(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	Delete(ctx context.Context, id int) error
}
