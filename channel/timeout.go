package main

import (
	"fmt"
	"time"
)

func main() {
	// Since select processes the first one received, if <-Time.After is faster, that process will run.
	// In order to use the select timeout pattern, you need to interact with the channel.
	chan1 := make(chan string, 1)
	chan2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		chan1 <- "one"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		chan2 <- "two"
	}()

	select {
	case msg1 := <-chan1:
		fmt.Println(msg1)
	case msg2 := <-chan1:
		fmt.Println(msg2)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
