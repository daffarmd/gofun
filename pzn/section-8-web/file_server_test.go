package section8web

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/daffarmd/gofun/helper"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fmt.Println(directory)
	fileServer := http.FileServer(directory)
	fmt.Println(fileServer)

	mux := http.NewServeMux()
	fmt.Println(mux)
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		helper.PanicErr(err)
	}
}
