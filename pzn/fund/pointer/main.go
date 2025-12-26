package main

import "fmt"

type Address struct {
	Kota      string
	Kecamatan string
	KodePos   int
}

func main() {
	addressKu := Address{"Malang", "Kromengan", 65165}
	addressChange := addressKu

	addressChange.KodePos = 77777

	fmt.Println(addressKu)
	fmt.Println(addressChange)
}
