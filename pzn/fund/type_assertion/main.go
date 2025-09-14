// tujuan nya adalah: memeriksa apakah suatu nilai dari tipe tertentu benar-benar merupakan tipe yang di harapkan.
package main

import "fmt"

func New() any {
	return "OK"
}

func main() {
	result := New()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("unknown type")
	}
}
