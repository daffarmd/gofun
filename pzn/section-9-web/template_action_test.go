package section8web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataIf(w http.ResponseWriter, r *http.Request) {

	dataMap := map[string]any{
		"Title": "Template",
		"Nama":  "Daffa",
	}

	t := template.Must(template.ParseFiles("./templates/if.html"))
	t.ExecuteTemplate(w, "if.html", dataMap)
}

func TemplateDataIfInt(w http.ResponseWriter, r *http.Request) {

	dataMap := map[string]any{
		"Nilai": 90,
		"Nama":  "Daffa",
	}

	t := template.Must(template.ParseFiles("./templates/ifcomparison.html"))
	t.ExecuteTemplate(w, "ifcomparison.html", dataMap)
}

func TemplateDataRange(w http.ResponseWriter, r *http.Request) {

	dataMap := map[string]any{
		"Hobbies": []string{
			"Gaming", "Coding", "Mancing", "Ngopi",
		},
	}

	t := template.Must(template.ParseFiles("./templates/range.html"))
	t.ExecuteTemplate(w, "range.html", dataMap)
}

func TemplateDataNested(w http.ResponseWriter, r *http.Request) {
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

	t := template.Must(template.ParseFiles("./templates/nested.html"))
	_ = t.ExecuteTemplate(w, "nested.html", data)
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	TemplateDataIf(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TestTemplateActionIfInt(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	TemplateDataIfInt(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	TemplateDataRange(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TestTemplateActionNestedData(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	TemplateDataNested(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}
