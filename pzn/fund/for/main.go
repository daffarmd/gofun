package main

import "fmt"

type siswa struct {
	id   int
	nama string
}

func loop(len int) {
	for i := 0; i < len; i++ {
		fmt.Println("Angka", i)
	}
}

func main() {
	loop(10)
}
