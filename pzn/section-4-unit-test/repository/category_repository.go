package repository

import "github.com/daffarmd/gofun/pzn/section-4-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
