package main

import "fmt"

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

// emebeded interface
type mockSuperman struct {
	Human // satisfy interface like this.
}

// Some methods can be overridden after satisfying the interface.
func (m mockSuperman) say() {
	fmt.Println("Overrided!")
}

func main() {
	m := &mockSuperman{}
	m.say()
}
