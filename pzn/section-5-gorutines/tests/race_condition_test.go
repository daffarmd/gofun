package tests

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				x = x + 1
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Count : ", x)
}
