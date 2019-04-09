package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("file/tmp/is_exist_file_or_directory.txt")

	// see https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go#answer-12527546
	if !os.IsNotExist(err) {
		fmt.Println("True")
		return
	}

	fmt.Println("False")
	return
}
