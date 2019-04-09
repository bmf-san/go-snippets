package main

import "fmt"

type handle func(s string)

func resolve(h handle) handle {
	return handle(h)
}

func handleS(s string) {
	fmt.Println(s)
}

func main() {
	s := "Hello World"

	resolve(handleS)(s)
}
