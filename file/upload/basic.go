package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	view := template.Must(template.ParseFiles("file/upload/index.html"))

	view.Execute(w, nil)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %v+\n", handler.Filename)
	fmt.Printf("Uploaded File: %v+\n", handler.Size)
	fmt.Printf("Uploaded File: %v+\n", handler.Header)

	tmpFile, err := ioutil.TempFile("file/upload/tmp", handler.Filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tmpFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpFile.Write(fileBytes)

	fmt.Fprintf(w, "Success uploaded file")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", uploadFile)

	http.ListenAndServe(":8080", nil)
}
