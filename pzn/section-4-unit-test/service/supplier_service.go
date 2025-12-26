package service

import (
	"errors"

	"github.com/daffarmd/gofun/pzn/section-4-unit-test/entity"
	"github.com/daffarmd/gofun/pzn/section-4-unit-test/repository"
)

type SupplierService struct {
	Repository repository.SupplierCategory
}

func (service SupplierService) Get(id int, kode_barang string) (*entity.Supplier, error) {
	supplier := service.Repository.FindById(id, kode_barang)
	if supplier == nil {
		return nil, errors.New("supplier not found")
	} else {
		return supplier, nil
	}
}
