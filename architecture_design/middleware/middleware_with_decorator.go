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
		fmt.Println("[start] middleware")
		next.ServeHTTP(w, r)
		fmt.Println("[end] middleware")
	}
}

func middlewareInside(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[start] middlewareInside")
		next.ServeHTTP(w, r)
		fmt.Println("[end] middlewareInside")
	}
}

type middleware func(http.HandlerFunc) http.HandlerFunc

type middlewareSet struct {
	middlewares []middleware
}

func newMiddlewareSet(mws ...middleware) middlewareSet {
	return middlewareSet{
		middlewares: append([]middleware(nil), mws...),
	}
}

func (mws middlewareSet) then(h http.HandlerFunc) http.HandlerFunc {
	for i := range mws.middlewares {
		h = mws.middlewares[len(mws.middlewares)-1-i](h)
	}

	return h
}

func main() {
	mws := newMiddlewareSet(middlewareOutside, middlewareInside)

	mux := http.NewServeMux()
	mux.HandleFunc("/", mws.then(index))

	http.ListenAndServe(":8080", mux)
}
