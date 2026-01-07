package tests

import (
	"fmt"
	"testing"
	"time"
)

func DisplayGoroutine(v int) {
	fmt.Println("Display : ", v)
}

func HelloGoroutine() {
	fmt.Println("Hello World")
}

func TestHelloGoroutine(t *testing.T) {
	go HelloGoroutine()
	fmt.Println("Test")

	time.Sleep(1 * time.Second)
}

func TestDisplayGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayGoroutine(i)
	}

	time.Sleep(3 * time.Second)

}
