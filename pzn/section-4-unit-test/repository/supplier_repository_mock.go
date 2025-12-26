package repository

import (
	"github.com/daffarmd/gofun/pzn/section-4-unit-test/entity"
	"github.com/stretchr/testify/mock"
)

type SupplierRepositoryMock struct {
	Mock mock.Mock
}

func (repository *SupplierRepositoryMock) FindById(id int, kode_barang string) *entity.Supplier {
	arguments := repository.Mock.Called(id, kode_barang)
	if arguments.Get(1) == nil {
		return nil
	} else {
		supplier := arguments.Get(0).(entity.Supplier)
		return &supplier
	}
}
