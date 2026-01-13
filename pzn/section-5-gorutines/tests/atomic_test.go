package tests

import (
	"fmt"
	"sync/atomic"
	"testing"
)

var counter1 int64 = 0

func TestAtomic(t *testing.T) {
	group.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			atomic.AddInt64(&counter1, 1)
		}
		group.Done()
	}()
	group.Wait()
	fmt.Println(counter1)
}
