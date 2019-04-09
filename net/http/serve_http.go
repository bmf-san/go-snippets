package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handler := new(HelloHandler)
	mux.Handle("/", handler)

	s := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	s.ListenAndServe()
}

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
