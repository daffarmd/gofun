package section8web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daffarmd/gofun/helper"
)

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.html", map[string]any{
		"Title": "Title",
		"Body":  "<p>Test<p>",
	})
}

func TestAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TestAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		helper.PanicErr(err)
	}
}

func TemplateAutoEscapeDisabled(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.html", map[string]any{
		"Title": "Title",
		"Body":  template.HTML("<p>Test<p>"),
	})
}

func TestAutoEscapeDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TestAutoEscapeServerDisabled(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	err := server.ListenAndServe()
	if err != nil {
		helper.PanicErr(err)
	}
}

func TemplateXss(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.html", map[string]any{
		"Title": "Title",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})
}

func TestXss(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000?body=<p>test</p", nil)
	recorder := httptest.NewRecorder()

	TemplateXss(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TestXssServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateXss),
	}

	err := server.ListenAndServe()
	if err != nil {
		helper.PanicErr(err)
	}
}

func TemplateXssEscape(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.html", map[string]any{
		"Title": "Title",
		"Body":  r.URL.Query().Get("body"),
	})
}

func TestXssEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateXssEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		helper.PanicErr(err)
	}
}
