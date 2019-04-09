package main

import (
	"fmt"
	"runtime"
	"sync"
)

var once = new(sync.Once)

func main() {
	// sync.Once allows the function to be executed only once.
	runtime.GOMAXPROCS(runtime.NumCPU())
	defer fmt.Println("Last executed")

	sw := new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		sw.Add(1)
		go func(sw *sync.WaitGroup) {
			once.Do(func() {
				fmt.Println("Run only once")
			})
			fmt.Println("Run continuously")
			sw.Done()
		}(sw)
	}
}
