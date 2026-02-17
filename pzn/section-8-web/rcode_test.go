package section8web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name is not valid")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hai %s", name)
	}
}

func Test(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000?name=Test", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
	fmt.Println(string(body))
}
