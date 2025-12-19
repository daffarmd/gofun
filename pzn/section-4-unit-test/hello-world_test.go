package section4unittest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Anis")
	assert.Equal(t, "Hello, Bambang!", result)

	fmt.Println("Lanjut ke kode ini")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Anis")
	require.Equal(t, "Hello, Bambang!", result)

	fmt.Println("Tidak lanjut ke kode ini")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Bambang")

	if result != "Hello, Bambang!" {
		panic("Result tidak sesuai")
	}

}

func TestHelloBakrie(t *testing.T) {
	result := HelloWorld("Bakrie")

	if result != "Hello, Bakrie!" {
		panic("Result tidak sesuai")
	}

}
