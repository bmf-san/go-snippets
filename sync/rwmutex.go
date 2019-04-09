package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex
var data map[int]string

func main() {
	// In addition to the Mutex function, RW Mutex can be locked to allow only Read.
	data = map[int]string{1: "A", 2: "B"}
	mu = sync.RWMutex{}
	go read()
	go read()
	go write()
	go read()
	time.Sleep(5 * time.Second)
}

func read() {
	mu.RLock()
	defer mu.RUnlock()
	time.Sleep(1 * time.Second)
	fmt.Println("read:", data[1])
}

func write() {
	mu.Lock()
	defer mu.Unlock()
	time.Sleep(2 * time.Second)
	data[1] = "Not A"
	fmt.Println("write", data[1])
}
