package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.OpenFile("file/tmp/open_file_with_option.txt", os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	fp.WriteString("Hello World")

	fp.Seek(0, 0)
	b := make([]byte, 216)
	n, _ := fp.Read(b)
	fmt.Println(string(b[:n]))
}
