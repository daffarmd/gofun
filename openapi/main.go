package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve file openapi.json langsung
	http.Handle("/api-spec-pdam-application.json", http.FileServer(http.Dir("./docs")))

	// Serve Swagger UI (static HTML + JS)
	// Kita pakai embed dari unpkg (CDN), jadi gak perlu install npm
	http.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		html := `
		<!DOCTYPE html>
		<html>
		<head>
		  <title>PDAM API Docs</title>
		  <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist/swagger-ui.css" />
		</head>
		<body>
		  <div id="swagger-ui"></div>
		  <script src="https://unpkg.com/swagger-ui-dist/swagger-ui-bundle.js"></script>
		  <script>
		    const ui = SwaggerUIBundle({
		      url: '/api-spec-pdam-application.json',
		      dom_id: '#swagger-ui',
		    });
		  </script>
		</body>
		</html>
		`
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
