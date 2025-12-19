package section4unittest

import "testing"

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
