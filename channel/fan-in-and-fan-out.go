package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producer(id int, out chan<- int) {
	for i := 0; i < 5; i++ {
		value := rand.Intn(100)
		fmt.Printf("Producer %d: Sending %d\n", id, value)
		out <- value
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	}
	close(out)
}

func fanIn(inputs []<-chan int, out chan<- int) {
	var wg sync.WaitGroup
	wg.Add(len(inputs))

	for _, input := range inputs {
		go func(ch <-chan int) {
			for value := range ch {
				out <- value
			}
			wg.Done()
		}(input)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Fan-Out
	numProducers := 3
	inputs := make([]chan int, numProducers)
	for i := 0; i < numProducers; i++ {
		inputs[i] = make(chan int)
		go producer(i+1, inputs[i])
	}

	// Convert channels to <-chan int
	inputChans := make([]<-chan int, numProducers)
	for i := 0; i < numProducers; i++ {
		inputChans[i] = inputs[i]
	}

	// Fan-In
	result := make(chan int)
	go fanIn(inputChans, result)

	// Consume the merged values
	for value := range result {
		fmt.Printf("Consumer: Received %d\n", value)
	}

	fmt.Println("All done!")
}
