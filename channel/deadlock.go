package main

import "fmt"

func main() {
	// fatal error: all goroutines are asleep - deadlock!
	msgs := make(chan string, 2)

	msgs <- "Hello"
	msgs <- "World"
	msgs <- "!"

	fmt.Println(<-msgs)
	fmt.Println(<-msgs)
	fmt.Println(<-msgs)
}
