package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`<!doctype html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>Go Server</title>
  </head>
  <body>
    <h1>Hello from Go</h1>
  </body>
</html>`))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handler)

	log.Println("listening on :8001")
	if err := http.ListenAndServe(":8001", mux); err != nil {
		log.Fatal(err)
	}
}
