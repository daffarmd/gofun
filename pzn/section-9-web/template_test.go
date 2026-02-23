package section8web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daffarmd/gofun/helper"
)

func HtmlTemplate(w http.ResponseWriter, r *http.Request) {
	templateText := "<html><body>{{.}}</body></html>"
	t, err := template.New("SIMPLE").Parse(templateText)
	if err != nil {
		helper.PanicErr(err)
	}

	t.ExecuteTemplate(w, "SIMPLE", "Hello HTML Template")
}

func TempletFileHtml(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/simple.html")
	if err != nil {
		helper.PanicErr(err)
	}

	t.ExecuteTemplate(w, "simple.html", "simple")
}

func TempletDirectoryHtml(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		helper.PanicErr(err)
	}

	t.ExecuteTemplate(w, "verysimple.html", "simple")
}

//go:embed templates/*.html
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(templates, "templates/*.html")
	if err != nil {
		helper.PanicErr(err)
	}

	t.ExecuteTemplate(w, "verysimple.html", "simple")
}

func TestTemplateHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	HtmlTemplate(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TestTemplateHtmlFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	TempletFileHtml(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TestDirectoryHtmlFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	TempletDirectoryHtml(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TestEmbedHtmlFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}
