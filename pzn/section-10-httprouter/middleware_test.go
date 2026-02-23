package section9httprouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type LogMiddleware struct {
	handler http.Handler
}

func (m *LogMiddleware) ServeHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Receive Request")
	m.handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Middleware Test")
	})

	middleware := LogMiddleware{handler: router}

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHttp(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Middleware Test", string(bytes))

	fmt.Println("end test")
}
