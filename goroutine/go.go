package main

import (
	"fmt"
	"time"
)

func say(v string) {
	for i := 0; i < 3; i++ {
		fmt.Println(v)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	// goroutine is a lightweight thread. By specifying a function in the go statement, it is executed in parallel.
	go say("goroutine")
	say("not goroutine")
	fmt.Println("done")
}
