package tests

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	totalThread := runtime.GOMAXPROCS(-1)
	totalGoroutine := runtime.NumGoroutine()

	fmt.Println(totalCPU)
	fmt.Println(totalThread)
	fmt.Println(totalGoroutine)

	group.Wait()
}
