package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var str bytes.Buffer

	log.Print("bytes buffer")

	fmt.Printf("%v", str.String()) // 2019/08/24 01:31:47 bytes buffer
}
