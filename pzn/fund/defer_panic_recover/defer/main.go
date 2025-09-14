package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func RunApplication() {
	defer logging()
	fmt.Println("Run Application")
}

func main() {
	RunApplication()
}
