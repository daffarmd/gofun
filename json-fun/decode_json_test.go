package jsonfun_test

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

func TestEncodeJson(t *testing.T) {
	jsonString := `{"FirstName": "Eko", "MiddleName": "asdad", "LastName": "Patrio"}`
	jsonEncode := []byte(jsonString)

	customer := &Customer{}

	fmt.Println("Customer : ", customer)

	json.Unmarshal(jsonEncode, &customer)

	fmt.Println("Customer Unmarshal : ", customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.MiddleName)

}
