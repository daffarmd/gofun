package section8web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000", nil)
	recorder := httptest.NewRecorder()

	HelloWorld(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	fmt.Println(body)
	if err != nil {
		panic(err)
	}

	bodyString := string(body)
	fmt.Println(bodyString)
}
