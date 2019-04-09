package main

import (
	"fmt"
	"time"
)

func main() {
	// It is used when you want to execute a certain part at a certain point in the future or at regular intervals.
	start := time.Now()
	timer := time.NewTimer(5 * time.Second)
	<-timer.C
	fmt.Println("Time!")
	end := time.Now()
	fmt.Printf("%f sec", (end.Sub(start)).Seconds())
}
