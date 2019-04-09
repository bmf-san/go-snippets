package main

import "fmt"

func sender(senders chan<- string, msg string) {
	senders <- msg
}

func receiver(senders <-chan string, receivers chan<- string) {
	msg := <-senders
	receivers <- msg
}

func main() {
	// You can use a channel as a function argument to specify whether you intend to send or receive.
	// "Hello" <- receivers "Hello" <- senders <- "Hello"
	senders := make(chan string, 1)
	receivers := make(chan string, 1)
	sender(senders, "Hello")
	receiver(senders, receivers)
	fmt.Println(<-receivers)
}
