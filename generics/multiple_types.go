package main

import "fmt"

func sum[T int | float64](nums []T) T {
	fmt.Printf("%T\n", nums)
	var sum T
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum[int]([]int{1, 2, 3, 4, 5}))                   // []int 15
	fmt.Println(sum[float64]([]float64{1.0, 2.0, 3.0, 4.0, 5.0})) // []float64 15
}
