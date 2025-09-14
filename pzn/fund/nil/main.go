package main

import "fmt"

func New(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	result := New("test")
	if result == nil {
		fmt.Println("Data not found")
	} else {
		fmt.Println(result)
	}
}
