package section8web

import (
	"net/http"
	"testing"

	"github.com/daffarmd/gofun/helper"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	getParam := r.URL.Query().Get("name")
	if getParam != "" {
		http.ServeFile(w, r, "./resources/ok.html")
	} else {
		http.ServeFile(w, r, "./resources/notfound.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		helper.PanicErr(err)
	}
}
