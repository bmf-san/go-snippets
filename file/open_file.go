package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readFileByOsRead() {
	fp, err := os.Open("file/tmp/open_file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	// Print a file path
	fmt.Println(fp.Name())

	b := make([]byte, 216)
	for {
		n, err := fp.Read(b)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(b[:n]))
	}

	return
}

func readFileByIoutilReadAll() {
	fp, err := os.Open("file/tmp/open_file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	// Print a file path
	fmt.Println(fp.Name())

	// To read the contents written most recently, you need to return to the top of the file with Seek.
	fp.Seek(0, 0)
	data, err := ioutil.ReadAll(fp)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(string(data))

	return
}

func main() {
	readFileByOsRead()
	readFileByIoutilReadAll()
}
