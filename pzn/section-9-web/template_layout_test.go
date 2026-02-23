package section8web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayouts(w http.ResponseWriter, r *http.Request) {

	dataMap := map[string]any{
		"Title": "Template",
		"Nama":  "Daffa",
	}

	t := template.Must(template.ParseFiles("./templates/header.html", "./templates/layouts.html", "./templates/footer.html"))
	t.ExecuteTemplate(w, "layouts.html", dataMap)
}

func TestTemplateLayouts(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	TemplateLayouts(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}
