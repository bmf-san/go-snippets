package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Create creates the specified file without returning an error even if it exists,
	// and deletes the contents of the original file.
	f, err := os.Create("file/tmp/create_file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	f.WriteString("Hello World")

	if err := f.Close; err != nil {
		fmt.Println(err)
		return
	}

	return
}
