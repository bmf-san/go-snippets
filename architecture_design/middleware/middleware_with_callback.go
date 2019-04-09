package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index")
}

func middlewareOutside(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[start] middlewareOutside")
		next.ServeHTTP(w, r)
		fmt.Println("[end] middlewareOutside")
	}
}

func middlewareInside(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[start] middlewareInside")
		next.ServeHTTP(w, r)
		fmt.Println("[end] middlewareInside")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", middlewareOutside(middlewareInside(index)))

	http.ListenAndServe(":8080", mux)
}
