package main

import "fmt"

func main() {
	// buffered channel can send values ​​up to a fixed amount without a corresponding receiver
	msgs := make(chan string, 2)

	msgs <- "Hello"
	msgs <- "World"

	fmt.Println(<-msgs)
	fmt.Println(<-msgs)
}
