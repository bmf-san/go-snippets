package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Post struct {
	ID   int
	Body string
}

func main() {
	src := `{"Id":3,"Body":"here is contents"}`

	p := &Post{}
	err := json.Unmarshal([]byte(src), p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", p) // &{ID:3 Body:here is contents}
}
