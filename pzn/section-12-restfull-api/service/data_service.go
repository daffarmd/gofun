package service

import (
	"context"

	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/model/web"
)

type DataService interface {
	Create(ctx context.Context, request web.DataCreateRequest) web.DataResponse
	Update(ctx context.Context, request web.DataUpdateRequest) web.DataResponse
	Delete(ctx context.Context, dataId int)
	FindById(ctx context.Context, dataId int) web.DataResponse
	FindAll(ctx context.Context) []web.DataResponse
	FindAllAsc(ctx context.Context) []web.DataResponse
}
