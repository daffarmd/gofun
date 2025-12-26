package service

import (
	"testing"

	"github.com/daffarmd/gofun/pzn/section-4-unit-test/entity"
	"github.com/daffarmd/gofun/pzn/section-4-unit-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var supplierRepository = &repository.SupplierRepositoryMock{Mock: mock.Mock{}}
var supplierService = SupplierService{Repository: supplierRepository}

func TestSupplierService_GetNotFound(t *testing.T) {

	// program mock
	supplierRepository.Mock.On("FindById", 1, "LPT11").Return(nil)

	category, err := supplierService.Get(1, "LPT11")
	assert.Nil(t, category)
	assert.NotNil(t, err)

}

func TestSupplierService_GetSuccess(t *testing.T) {
	supplier := entity.Supplier{
		Id:         2,
		NamaBarang: "Asus",
		KodeBarang: "LPT11",
	}
	supplierRepository.Mock.On("FindById", 2, "LPT11").Return(supplier, nil)

	result, err := supplierService.Get(2, "LPT11")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, supplier.Id, result.Id)
	assert.Equal(t, supplier.NamaBarang, result.NamaBarang)
	assert.Equal(t, supplier.KodeBarang, result.KodeBarang)
}
