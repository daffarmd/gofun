package section8web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PageData struct {
	Title  string
	User   UserData
	Alamat AddressData
}

type UserData struct {
	Nama  string
	Email string
}

type AddressData struct {
	Kota   string
	Negara string
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Template",
		User: UserData{
			Nama:  "Daffa",
			Email: "daffa@mail.com",
		},
		Alamat: AddressData{
			Kota:   "Jakarta",
			Negara: "Indonesia",
		},
	}

	t := template.Must(template.ParseFiles("./templates/name_struct.html"))
	_ = t.ExecuteTemplate(w, "name_struct.html", data)
}

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {

	dataMap := map[string]any{
		"Title": "Template",
		"Nama":  "Daffa",
	}

	t := template.Must(template.ParseFiles("./templates/name.html"))
	t.ExecuteTemplate(w, "name.html", dataMap)
}

func TestTemplateData(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}
