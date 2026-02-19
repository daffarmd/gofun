package section8web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var myTemplates = template.Must(template.ParseFS(templates, "templates/*.html"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "simple.html", "Hello Templates")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}
