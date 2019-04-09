package main

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	Value int
}

func main() {
	// sync.Pool is a thread-safe memory pool.
	p := sync.Pool{
		New: func() interface{} {
			return &Data{
				Value: 0,
			}
		},
	}
	sw := &sync.WaitGroup{}

	pf := func() {
		d := p.Get().(*Data)
		if d.Value == 0 {
			fmt.Println("New")
		} else {
			fmt.Println("Cache")
		}
		d.Value++
		p.Put(d)
		time.Sleep(1 * time.Microsecond)
		sw.Done()
	}

	for i := 0; i < 10; i++ {
		sw.Add(1)
		go pf()
	}

	for i := 0; i < 10; i++ {
		sw.Add(1)
		go pf()
	}

	sw.Wait()
	fmt.Println("End")
}
