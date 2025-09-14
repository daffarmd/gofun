package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("terjadi panic : ", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error karena gaji kurang")
	}
}

func main() {
	runApp(true)
	fmt.Println("Jika tidak di recover maka ketika terjadi panic code program ini tidak akan di eksekusi")
}
