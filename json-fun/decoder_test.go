package jsonfun_test

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"

	"os"
	"testing"
)

func TestDecoder(t *testing.T) {
	readfile, _ := os.Open("inijson.json")
	decoder := json.NewDecoder(readfile)
	customer := Customer{}
	_ = decoder.Decode(&customer)

	fmt.Println(customer)
}
