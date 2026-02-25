package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Handler untuk mengambil data JSON
func GetTranslationHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Tentukan path ke file JSON
	filePath := "assets/translation/indonesia_translate.json"

	// 2. Buka file
	jsonFile, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File tidak ditemukan", http.StatusNotFound)
		return
	}
	defer jsonFile.Close()

	// 3. Baca isi file
	byteValue, _ := io.ReadAll(jsonFile)

	// 4. Set Header agar browser/frontend tahu ini adalah JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// 5. Kirim isi file sebagai response
	w.Write(byteValue)
}

func main() {
	http.HandleFunc("/api/translations", GetTranslationHandler)

	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
