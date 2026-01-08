package tests

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce) // hanya satu yang jalan
			// OnlyOnce() // berjalan semua
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Done", counter)
}
