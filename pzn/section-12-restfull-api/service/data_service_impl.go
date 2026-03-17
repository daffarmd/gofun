package service

import (
	"context"
	"database/sql"

	"github.com/daffarmd/gofun/helper"
	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/exception"
	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/model/domain"
	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/model/web"
	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/repository"
	"github.com/go-playground/validator"
)

type DataServiceImpl struct {
	DataRepository repository.DataRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewDataService(dr repository.DataRepository, DB *sql.DB, v *validator.Validate) DataService {
	return &DataServiceImpl{
		DataRepository: dr,
		DB:             DB,
		Validate:       v,
	}
}

func (service *DataServiceImpl) Create(ctx context.Context, request web.DataCreateRequest) web.DataResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	data := domain.Data{
		Name: request.Name,
	}

	data = service.DataRepository.Create(ctx, tx, data)

	return helper.ToDataResponse(data)

}

func (service *DataServiceImpl) Update(ctx context.Context, request web.DataUpdateRequest) web.DataResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	data, err := service.DataRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	data.Name = request.Name

	data = service.DataRepository.Update(ctx, tx, data)

	return helper.ToDataResponse(data)
}

func (service *DataServiceImpl) Delete(ctx context.Context, dataId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	data, err := service.DataRepository.FindById(ctx, tx, dataId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.DataRepository.Delete(ctx, tx, data)
}

func (service *DataServiceImpl) FindById(ctx context.Context, dataId int) web.DataResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	data, err := service.DataRepository.FindById(ctx, tx, dataId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToDataResponse(data)
}

func (service *DataServiceImpl) FindAll(ctx context.Context) []web.DataResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	data := service.DataRepository.FindAll(ctx, tx)

	return helper.ToDataResponses(data)
}

func (service *DataServiceImpl) FindAllAsc(ctx context.Context) []web.DataResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	data := service.DataRepository.FindAllAsc(ctx, tx)

	return helper.ToDataResponses(data)
}
