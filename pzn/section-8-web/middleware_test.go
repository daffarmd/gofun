package section8web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

type ErrorHandler struct {
	Handler http.Handler
}

func (m *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before Middleware")
	m.Handler.ServeHTTP(w, r)
	fmt.Println("After Middleware")
}

func (m *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		fmt.Println("RECOVER : ", err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error : %v", err)
		}
	}()
	m.Handler.ServeHTTP(w, r)
}

func TestHandlerMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Execute Middleware")
		fmt.Fprintf(w, "Hello Middleware")
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Server Internal Error")
	})

	logMiddleware := &LogMiddleware{Handler: mux}
	errorHandler := &ErrorHandler{Handler: logMiddleware}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
