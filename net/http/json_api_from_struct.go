package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Items Items  `json:"items"`
}

type Items struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:   1,
		Name: "foo",
		Items: Items{
			ID:    1,
			Title: "Hello world",
		},
	}

	json.NewEncoder(w).Encode(user)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	http.ListenAndServe(":8080", mux)
	// {"id":1,"name":"foo","items":{"id":1,"title":"Hello world"}}
}
