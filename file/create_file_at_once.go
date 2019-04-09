package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data := "Hello World"

	err := ioutil.WriteFile("file/tmp/create_file_at_once.txt", []byte(data), 0664)
	if err != nil {
		fmt.Println(err)
	}
}
