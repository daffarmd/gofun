package jsonfun_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TesetEncoder(t *testing.T) {
	writeFile, _ := os.Create("out.json")
	encoder := json.NewEncoder(writeFile)

	cust := Customer{
		FirstName:  "Budi",
		MiddleName: "ono",
		LastName:   "Siregar",
	}

	encoder.Encode(cust)

	fmt.Println(cust)
}
