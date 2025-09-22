package services

import (
	"context"
	"database/sql"

	"github.com/daffarmd/gofun/pzn/section-12/helper"
	"github.com/daffarmd/gofun/pzn/section-12/model/domain"
	"github.com/daffarmd/gofun/pzn/section-12/model/web"
	"github.com/daffarmd/gofun/pzn/section-12/repository"
	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func (service *CategoryServiceImpl) Save(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToModeCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryCreateRequestUpdate) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	category := domain.Category{
		Id:   request.Id,
		Name: request.Name,
	}

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToModeCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToModeCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context, request web.CategoryCreateRequest) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
