package section8web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My name is " + myPage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Bambang" }}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Arip",
	})
}

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ len .Name }}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Arip",
	})
}

func TemplateFunctionMap(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]any{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse("{{ upper .Name}}"))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Arip",
	})
}

func TemplateFunctionPipelines(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]any{
		"sayHello": func(value string) string {
			return "Hello " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse("{{ sayHello .Name | upper}}"))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Arip",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TestTemplateFunctionMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionMap(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TestTemplateFunctionPipelines(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionPipelines(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}
