package main

import (
	"fmt"
	"time"
)

func main() {
	// Ticker is used when doing something at regular intervals.
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Stopped")
}
