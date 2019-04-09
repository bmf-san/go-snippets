package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// sync.Mutex a function to prevent errors due to conflicts.
	// Lock using Mutex can prohibit both read (Read) and write (Write).
	runtime.GOMAXPROCS(runtime.NumCPU())

	sw := new(sync.WaitGroup)
	m := new(sync.Mutex)
	for i := 0; i < 10; i++ {
		sw.Add(1)
		go func(sw *sync.WaitGroup, m *sync.Mutex) {
			m.Lock()
			defer m.Unlock()
			fmt.Println("A")
			time.Sleep(100 * time.Millisecond)
			fmt.Println("B")
			time.Sleep(100 * time.Millisecond)
			fmt.Println("C")
			time.Sleep(100 * time.Millisecond)
			sw.Done()
		}(sw, m)
		sw.Wait()
	}
}
