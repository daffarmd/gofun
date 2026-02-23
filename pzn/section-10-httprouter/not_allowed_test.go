package section9httprouter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestNotAllowed(t *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Not Allowed")
	})
	router.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Post")
	})

	request := httptest.NewRequest("GET", "http://127.0.0.1:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	fmt.Println(recorder.Code)
	fmt.Println(recorder.Body.String())
	fmt.Println(recorder.Header().Get("Allow"))

	assert.Equal(t, http.StatusMethodNotAllowed, recorder.Code)
	assert.Equal(t, "Not Allowed", recorder.Body.String())
	assert.Equal(t, "POST", recorder.Header().Get("Allow"))
}
