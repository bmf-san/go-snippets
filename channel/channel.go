package main

import "fmt"

func main() {
	msgs := make(chan string)
	go func() {
		msgs <- "Hello"
	}()

	msg := <-msgs
	fmt.Println(msg)
}
