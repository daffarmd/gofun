package helper

import (
	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/model/domain"
	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/model/web"
)

func ToDataResponse(data domain.Data) web.DataResponse {
	return web.DataResponse{
		Id:     data.Id,
		Name:   data.Name,
		Status: data.Status,
	}
}

func ToDataResponses(data []domain.Data) []web.DataResponse {
	var dataResponses []web.DataResponse
	for _, datas := range data {
		dataResponses = append(dataResponses, ToDataResponse(datas))
	}

	return dataResponses
}
