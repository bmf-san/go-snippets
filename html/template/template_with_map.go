package main

import (
	"html/template"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("html/template/sample.tpl"))

	m := map[string]string{
		"Date": time.Now().Format("2006-01-02"),
		"Name": "foo",
	}

	tpl.Execute(w, m)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	http.ListenAndServe(":8080", mux)
}
