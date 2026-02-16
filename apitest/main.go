package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Struct untuk request
type NameRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Struct untuk response
type NameResponse struct {
	FullName string `json:"fullname"`
}

func main() {
	http.HandleFunc("/fullname", handleFullName)

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleFullName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req NameRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.FirstName == "" || req.LastName == "" {
		http.Error(w, "firstname and lastname are required", http.StatusBadRequest)
		return
	}

	response := NameResponse{
		FullName: req.FirstName + " " + req.LastName,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
