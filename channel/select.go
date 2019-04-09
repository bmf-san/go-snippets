package main

import (
	"fmt"
	"time"
)

func main() {
	// By using select, you can wait for multiple channel operations.
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "channel 2"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		chan1 <- "one"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		}
	}

}
