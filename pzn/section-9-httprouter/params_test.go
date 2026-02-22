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

func ParamsRouter(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	text := "Product " + params.ByName("id")
	fmt.Fprint(w, text)
}
func TestRouterParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", ParamsRouter)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
	assert.Equal(t, "Product 2", string(bytes))
}
