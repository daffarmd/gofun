package section8web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/daffarmd/gofun/helper"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	err := myTemplates.ExecuteTemplate(w, "upload-form.html", nil)
	if err != nil {
		panic(err)
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload-form-success.html", map[string]any{
		"Title": "Success",
		"Name":  name,
		"file":  "/static/" + fileHeader.Filename,
	})

}

func TestUploadFormServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./resources"))))
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		helper.PanicErr(err)
	}
}

//go:embed resources/32.jpg
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	w.WriteField("name", "Certificate")

	file, _ := w.CreateFormFile("file", "test-upload.jpg")
	file.Write(uploadFileTest)
	w.Close()

	r := httptest.NewRequest(http.MethodPost, "http://localhost:7070", body)
	r.Header.Set("Content-Type", w.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, r)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
