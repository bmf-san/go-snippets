package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	b, err := os.Open("file/tmp/scan_file.txt")
	if err != nil {
		fmt.Println(err)
	}
	s := bufio.NewScanner(b)

	// if you don't need a line, you can omit a { and i.
	for i := 1; s.Scan(); i++ {
		line := s.Text()
		fmt.Println(i, line)
	}

	if err := s.Err(); err != nil {
		fmt.Println(err)
	}
}
