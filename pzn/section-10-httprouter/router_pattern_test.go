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

func RouterPattern(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	text := "Product " + params.ByName("id") + " item " + params.ByName("itemId")
	fmt.Fprint(w, text)
}

func TestRouterPattern(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/item/:itemId", RouterPattern)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/2/item/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	fmt.Println("test")
	fmt.Println(string(bytes))
	assert.Equal(t, "Product 2 item 2", string(bytes))
}

func CatchAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	text := "Image : " + params.ByName("image")
	fmt.Fprint(w, text)
}

func TestCatchAll(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", CatchAll)

	// route := "http://localhost:8080/images/small/small.png"
	route2 := "http://localhost:8080/images/image.jpg"

	request := httptest.NewRequest(http.MethodGet, route2, nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	fmt.Println("test")
	fmt.Println(string(bytes))
	assert.Equal(t, "Image : /image.jpg", string(bytes))
}
