package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("GET"))
	case "POST":
		w.Write([]byte("POST"))
	case "PUT":
		w.Write([]byte("PUT"))
	case "PATCH":
		w.Write([]byte("PATCH"))
	case "DELETE":
		w.Write([]byte("DELETE"))
	}
}

func middlewareCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Access-Control-Allow-Origin", "http://localhost:1234")
		r.Header.Set("Access-Control-Max-Age", "86400")
		r.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		r.Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		r.Header.Set("Access-Control-Expose-Headers", "Content-Length")
		r.Header.Set("Access-Control-Allow-Credentials", "true")

		w.WriteHeader(http.StatusOK)

		next.ServeHTTP(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", middlewareCORS(index))

	if err := http.ListenAndServe(":1234", mux); err != nil {
		fmt.Println(err)
	}
}
