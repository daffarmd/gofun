package main

import (
	"fmt"

	one "github.com/daffarmd/gofun/play-ground/function_task/task_function_as_parameter"
)

func main() {
	// NO : 1

	p := []one.Product{
		{Name: "Laptop XYZ", Category: "Elektronik", Price: 15000000},
		{Name: "Baju Koko", Category: "Fashion", Price: 250000},
		{Name: "Headphone ABC", Category: "Elektronik", Price: 800000},
		{Name: "Sepatu Sport", Category: "Fashion", Price: 1200000},
	}

	filterProduct := one.FilterProducts(p, func(p one.Product) bool {
		return p.Name == "Laptop XYZ"
	})

	fmt.Println("Produk filtered : ", filterProduct)

	// NO 2
	harga := []int{10000, 25000, 40000, 55000}
	potonganBerapa := func(harga int) int {
		if harga < 30000 {
			return harga * 10 / 100
		} else {
			return harga * 20 / 100
		}
	}
	diskon := func(harga int) int {
		if harga < 30000 {
			return harga - (harga * 10 / 100)
		} else {
			return harga - (harga * 20 / 100)
		}
	}
	for _, v := range harga {
		fmt.Printf("Harga Awal : %d | Potongan : %d | Harga Discount : %d \n", v, potonganBerapa(v), diskon(v))
	}
}
