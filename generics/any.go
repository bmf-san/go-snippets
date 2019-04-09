package main

import "fmt"

// any is an alias for interface{}
func print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	print[int]([]int{1, 2, 3, 4, 5})
	print[string]([]string{"a", "b", "c", "d", "e"})
}
