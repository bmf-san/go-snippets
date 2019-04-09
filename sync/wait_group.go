package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// sync.WaitGroup can wait for the completion of multiple Goroutines.
	var sw sync.WaitGroup
	for i := 0; i < 10; i++ {
		sw.Add(1)
		go func(i int) {
			time.Sleep(2 * time.Second)
			fmt.Println(i)
			sw.Done()
		}(i)
	}
	sw.Wait()
}
