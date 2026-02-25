package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTranslationHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/translations", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTranslationHandler)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler mengembalikan status code salah: dapat %v ingin %v",
			status, http.StatusOK)
	}

	expectedHeader := "application/json"
	if contentType := recorder.Header().Get("Content-Type"); contentType != expectedHeader {
		t.Errorf("handler mengembalikan content type salah: dapat %v ingin %v",
			contentType, expectedHeader)
	}

	if recorder.Body.Len() == 0 {
		t.Errorf("handler mengembalikan body kosong")
	}

	fmt.Println(recorder.Body)
}
