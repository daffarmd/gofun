package repository

import (
	"context"
	"database/sql"

	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/model/domain"
)

type DataRepository interface {
	Create(ctx context.Context, tx *sql.Tx, data domain.Data) domain.Data
	Update(ctx context.Context, tx *sql.Tx, data domain.Data) domain.Data
	Delete(ctx context.Context, tx *sql.Tx, data domain.Data)
	FindById(ctx context.Context, tx *sql.Tx, data int) (domain.Data, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Data
	FindAllAsc(ctx context.Context, tx *sql.Tx) []domain.Data
}
