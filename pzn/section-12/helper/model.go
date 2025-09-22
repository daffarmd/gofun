package helper

import (
	"github.com/daffarmd/gofun/pzn/section-12/model/domain"
	"github.com/daffarmd/gofun/pzn/section-12/model/web"
)

func ToModeCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}

}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToModeCategoryResponse(category))
	}
	return categoryResponses
}
