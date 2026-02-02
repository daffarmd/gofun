package section6context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	bg := context.Background()
	bgTodo := context.TODO()

	fmt.Println(bg)
	fmt.Println(bgTodo)

}

func TestContextWithValue(t *testing.T) {
	ctxA := context.Background()

	ctxB := context.WithValue(ctxA, "b", "B")
	ctxC := context.WithValue(ctxA, "c", "C")

	ctxD := context.WithValue(ctxB, "d", "D")
	ctxE := context.WithValue(ctxB, "e", "E")

	ctxF := context.WithValue(ctxC, "f", "F")

	fmt.Println("Context Parent :", ctxA)
	fmt.Println("Context Child B :", ctxB)
	fmt.Println("Context Child C :", ctxC)

	fmt.Println("Context Child D :", ctxD)
	fmt.Println("Context Child E :", ctxE)
	fmt.Println("Context Child F :", ctxF)

	// GET VALUE
	fmt.Println("Get Value Context D:", ctxD.Value("c"))

}

// GOROUTINE LEAK
func CreateCounter() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			destination <- counter
			counter++
		}
	}()

	return destination
}

func CreateCounterCancel(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func TestContextWithCancelLeak(t *testing.T) {
	fmt.Println("Jumlah Goroutine Awal: ", runtime.NumGoroutine())

	destination := CreateCounter()
	for n := range destination {
		fmt.Println("Counter : ", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Jumlah Goroutine Akhir: ", runtime.NumGoroutine())
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Jumlah Goroutine Awal: ", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounterCancel(ctx)
	for n := range destination {
		fmt.Println("Counter : ", n)
		if n == 10 {
			break
		}
	}
	cancel()

	time.Sleep(1 * time.Second) // menunggu semua proses selesai

	fmt.Println("Jumlah Goroutine Akhir: ", runtime.NumGoroutine())
}

func CreateCounterTimeout(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // simulasi slow
			}
		}
	}()

	return destination
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Jumlah Goroutine Awal: ", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 10*time.Second)
	defer cancel()

	destination := CreateCounterTimeout(ctx)
	for n := range destination {
		fmt.Println("Counter : ", n) // looping unlimited
	}

	time.Sleep(1 * time.Second) // menunggu semua proses selesai

	fmt.Println("Jumlah Goroutine Akhir: ", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Jumlah Goroutine Awal: ", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounterTimeout(ctx)
	for n := range destination {
		fmt.Println("Counter : ", n) // looping unlimited
	}

	time.Sleep(1 * time.Second) // menunggu semua proses selesai

	fmt.Println("Jumlah Goroutine Akhir: ", runtime.NumGoroutine())
}
