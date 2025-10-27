// Array dalam json di representasikan dengan [] atau Slice dalam bahasa golang
// Format String bisa pakai Sprintf atau Printf
// Sprintf dia untuk disimpan di variable
// Printf dia untuk langsung di print ke log atau termial
package jsonfun_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Addres struct {
	City     string
	PostCode string
}

type DataPegawai struct {
	Nama   string
	Nip    string
	Hobbi  []string
	Age    int
	Addres []Addres
}

func TestEncodeArray(t *testing.T) {

	data := []DataPegawai{
		{
			Nama:  "Eko",
			Nip:   "123",
			Hobbi: []string{"Tidur", "Ngoding"},
			Age:   12,
		},
	}

	b, _ := json.Marshal(data)

	fmt.Println(string(b))
}

func TestDecodeArray(t *testing.T) {

	jsonString := `[{"Nama": "Eko", "Nip": "123", "Hobbi": ["Tidur", "Ngoding"]}]`

	jsonDecode := []byte(jsonString)

	var data []DataPegawai

	json.Unmarshal(jsonDecode, &data)

	fmt.Println(data)

}

func TestArrayComplex(t *testing.T) {
	dataArrCom := []DataPegawai{
		{
			Nama: "Eko",
			Nip:  "123123",
			Age:  12,
			Hobbi: []string{
				"tidur",
				"ngoding",
			},
			Addres: []Addres{
				{
					City:     "Jakarta",
					PostCode: "123",
				},
			},
		},
	}

	b, _ := json.Marshal(dataArrCom)
	fmt.Println(string(b))
}

func TestOnlyArrayComplex(t *testing.T) {
	dataArrCom := []Addres{
		{
			City:     "Jakarta",
			PostCode: "123",
		},
	}

	b, _ := json.Marshal(dataArrCom)
	fmt.Println(string(b))
}

func TestArrayDecodeComplex(t *testing.T) {
	dataReq := `[{"Nama":"Eko","Nip":"123123","Hobbi":["tidur","ngoding"],"Age":12,"Addres":[{"City":"Jakarta","PostCode":"123"}]}]`
	jsonBytes := []byte(dataReq)
	var dataArrCom []DataPegawai
	json.Unmarshal(jsonBytes, &dataArrCom)
	fmt.Println(dataArrCom)
}

func TestOnlyArrayDecodeComplex(t *testing.T) {
	dataReq := `[{"City":"Jakarta","PostCode":"123"}]`
	jsonBytes := []byte(dataReq)
	var dataArrCom []Addres
	json.Unmarshal(jsonBytes, &dataArrCom)
	fmt.Println(dataArrCom)
}
