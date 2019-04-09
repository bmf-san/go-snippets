package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

func index(w http.ResponseWriter, r *http.Request) {
	view := template.Must(template.ParseFiles("file/upload/index.html"))

	view.Execute(w, nil)
}

const maxUploadSize = 2 * 1024 * 1024

func uploadFile(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		fmt.Println("File too big")
		fmt.Println(err)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("invalid file")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("invalid error")
		fmt.Println(err)
		return
	}

	fileType := http.DetectContentType(fileBytes)
	switch fileType {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/pdf":
		break
	default:
		fmt.Println("invalid file type")
		fmt.Println(err)
		return
	}

	ext, err := mime.ExtensionsByType(fileType)
	if err != nil {
		fmt.Println("invalid file")
		fmt.Println(err)
		return
	}

	fileName := "upload" + ext[0]

	newFile, err := os.Create(filepath.Join("file/upload/tmp", fileName))
	if err != nil {
		fmt.Println("invalid file")
		fmt.Println(err)
		return
	}
	defer newFile.Close()

	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		fmt.Println("can't write file")
		fmt.Println(err)
		return
	}
	w.Write([]byte("Success"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", uploadFile)

	http.ListenAndServe(":8080", nil)
}
