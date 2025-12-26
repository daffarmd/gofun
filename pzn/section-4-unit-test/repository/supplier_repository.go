package repository

import "github.com/daffarmd/gofun/pzn/section-4-unit-test/entity"

type SupplierCategory interface {
	FindById(id int, kode_barang string) *entity.Supplier
}
