package services

import (
	"context"

	"github.com/daffarmd/gofun/pzn/section-12/model/web"
)

type CategoryService interface {
	Save(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryCreateRequestUpdate) web.CategoryResponse
	Delete(ctx context.Context, categoryID int)
	FindById(ctx context.Context, categoryID int)
	FindAll(ctx context.Context) []web.CategoryResponse
}
