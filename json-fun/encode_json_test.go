package jsonfun_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("Test")
	logJson([]string{"Eko", "Patrio"})
}
