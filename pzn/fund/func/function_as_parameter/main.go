package main

import "fmt"

// version 01
// func sayHelloWithFilter(nama string, filter func(string) string) {
// 	filteredName := filter(nama)
// 	fmt.Println("Hallo", filteredName)
// }

type Filter func(string) string

func sayHelloWithFilter(nama string, filter Filter) {
	filteredName := filter(nama)
	fmt.Println("Hallo", filteredName)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "******"
	} else {
		return name
	}
}

func main() {
	filter := spamFilter

	sayHelloWithFilter("test", filter)
	sayHelloWithFilter("anjing", filter)
}
