package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	stop := make(chan bool)
	go func() {
	loop:
		for {
			select {
			case t := <-ticker.C:
				fmt.Println(t)
			case <-stop:
				break loop
			}
		}
		fmt.Println("ok")
	}()

	time.Sleep(3 * time.Second)
	ticker.Stop()
	close(stop)
}
