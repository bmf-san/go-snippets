package main

import "fmt"

func defaultArgument(s ...string) {
	if len(s) == 0 {
		fmt.Println("no arguments.")
		return
	}
	fmt.Println("argument")
	return
}

func main() {
	defaultArgument("foo")
	defaultArgument()
}
