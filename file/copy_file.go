package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	w, err := os.Create("file/tmp/original.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer w.Close()

	r, err := os.Open("file/tmp/copy_file.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer r.Close()

	_, err = io.Copy(w, r)
	if err != nil {
		fmt.Println(err)
	}
}
