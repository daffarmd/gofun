package section9httprouter

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	dir, _ := fs.Sub(resources, "resources")

	router.ServeFiles("/resources/*filepath", http.FS(dir))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/resources/ok.html", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
