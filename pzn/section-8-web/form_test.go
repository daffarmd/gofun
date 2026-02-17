package section8web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/daffarmd/gofun/helper"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	helper.HttpError(err, w)

	firstname := r.PostForm.Get("first_name")
	lastname := r.PostForm.Get("last_name")
	fmt.Fprintf(w, "Hello %s %s", firstname, lastname)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Daf&last_name=Devs")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:7000", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(body)
	fmt.Println(bodyString)
}
