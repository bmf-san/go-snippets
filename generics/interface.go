package main

import "fmt"

type Numeric interface {
	int | float64
}

func print[T Numeric](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	print[int]([]int{1, 2, 3, 4, 5})
	print[float64]([]float64{1.0, 2.0, 3.0, 4.0, 5.0})
}
