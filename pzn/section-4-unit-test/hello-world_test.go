package section4unittest

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Anis",
			request:  "Anis",
			expected: "Hello, Anis!",
		},
		{
			name:     "Baswedan",
			request:  "Baswedan",
			expected: "Hello, Baswedan!",
		},
		{
			name:     "Presiden Aceh",
			request:  "Presiden Aceh",
			expected: "Hello, Presiden Aceh!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.name)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		require.Equal(t, "Hello, Eko!", result)
	})

	t.Run("Kurniawan", func(t *testing.T) {
		result := HelloWorld("Kurniawan")
		require.Equal(t, "Hello, Kurniawan!", result)
	})
}

func TestMain(t *testing.M) {
	fmt.Println("sebelum unit test")

	t.Run()

	fmt.Println("sesudah unit test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Os tidak cocok")
	}

	result := HelloWorld("Bakrie")
	assert.Equal(t, "Hello, Bambang!", result)
}

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
