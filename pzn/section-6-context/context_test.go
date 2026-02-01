package section6context

import (
	"context"
	"fmt"
	"testing"
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
