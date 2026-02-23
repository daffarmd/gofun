package section11json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func EncodeJson(data any) {
	bytes, _ := json.Marshal(data)

	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	var person = map[string]any{
		"Nama": []string{
			"Budi",
			"Bambang",
			"Sushi",
		},
	}
	EncodeJson(person)

	fmt.Println(person["Nama"].([]string)[1])

	// Out : Nama :  [Budi Bambang Sushi]
	for _, persons := range person {
		fmt.Println("Nama : ", persons)
	}

	// Out :
	// Nama :  Budi
	// Nama :  Bambang
	// Nama :  Sushi
	names := person["Nama"].([]string)
	for _, persons := range names {
		fmt.Println("Nama : ", persons)
	}
}

func TestAngon(t *testing.T) {
	payload := map[string]any{
		"data": []any{
			1388,
			"aV6cmoWCrIXSdZOZhnKbxFGhoqpspZ%2drfo_OaWpyU16Bv3aoYnKRmHZVq2fLp5mjdIGrsXl%2dqZVmXaVxk1yQqGlfnJmGqw",
			"1771811982",
			1,
			1,
			`[["test"]]`,
		},
	}

	bytes, _ := json.Marshal(payload)
	fmt.Println(string(bytes))
}

func TestAngonDecode(t *testing.T) {
	payload := `{"data":[1388,"aV6cmoWCrIXSdZOZhnKbxFGhoqpspZ%2drfo_OaWpyU16Bv3aoYnKRmHZVq2fLp5mjdIGrsXl%2dqZVmXaVxk1yQqGlfnJmGqw","1771811982",1,1,"[[\"test\"]]"]}`

	var result map[string]any

	err := json.Unmarshal([]byte(payload), &result)
	if err != nil {
		t.Fatal(err)
	}

	data := result["data"].([]any)

	fmt.Println("ID:", data[0])
	fmt.Println("Token:", data[1])
	fmt.Println("Timestamp:", data[2])
	fmt.Println("Flag1:", data[3])
	fmt.Println("Flag2:", data[4])
	// Response Tidak Ideal
	fmt.Println("Raw Nested:", data[5])
}
