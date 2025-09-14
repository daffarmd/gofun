package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(v int) int {
	if v == 1 {
		return 1
	} else {
		return v * factorialRecursive(v-1)
	}
}

func main() {
	fmt.Println(factorialRecursive(10))

	fmt.Println(factorialLoop(10))
}
