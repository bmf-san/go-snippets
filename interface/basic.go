package main

import "fmt"

// All methods need to be implemented to satisfy interface.
type Human interface {
	say()
	walk()
}

type Superman struct{}

func (s *Superman) say() {
	return
}

func (s *Superman) walk() {
	return
}

func (s *Superman) fly() {
	return
}

func main() {
	s := &Superman{}
	i, ok := interface{}(s).(Human)
	if !ok {
		fmt.Print("Not implemented Human")
	}
	fmt.Println(i, ok) // {} true
}
