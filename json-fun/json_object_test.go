package jsonfun

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Eko",
		MiddleName: "Patrio",
		LastName:   "Kurniawan",
	}
	fmt.Println(customer)

	b, _ := json.Marshal(customer)
	fmt.Println(b)
	fmt.Println(string(b))

	jsonRequest := `{
    "FirstName": "Eko",
    "MiddleName": "Patrio",
    "LastName": "Kurniawan"
}`
	jsonBytes := []byte(jsonRequest)
	var customer2 Customer
	json.Unmarshal(jsonBytes, &customer2)
	fmt.Println(customer2)

}

// Response = Marshal
// Request = Unmarshal
