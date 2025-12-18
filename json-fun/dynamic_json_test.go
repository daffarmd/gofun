package jsonfun_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDynamicMapDecode(t *testing.T) {
	jsonString := `{"name_user":"bambang","umur":12}`
	jsonBytes := []byte(jsonString)

	var mapStr map[string]interface{}
	json.Unmarshal(jsonBytes, &mapStr)

	fmt.Println(mapStr)
	fmt.Println(mapStr["umur"])

}
