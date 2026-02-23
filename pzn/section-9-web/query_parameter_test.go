package section8web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000/?name=Sapri", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(body)
	fmt.Println(bodyString)
}

func MultipleParam(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	alamat := r.URL.Query().Get("alamat")

	if name == "" || alamat == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s from %s", name, alamat)
	}
}

func TestQueryMultipelParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000/?name=Sapri&alamat=Indonesia", nil)
	recorder := httptest.NewRecorder()

	MultipleParam(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(body)
	fmt.Println(bodyString)
}

func MultipleValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	nameJoin := strings.Join(names, " ")
	fmt.Fprint(w, "Hello ", nameJoin)
}

func TestQueryMultipelValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000/?name=Sapri&name=Hartono", nil)
	recorder := httptest.NewRecorder()

	MultipleValues(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(body)
	fmt.Println(bodyString)
}
