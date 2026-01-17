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
