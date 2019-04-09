package main

import (
	"log"
	"reflect"
)

// diff compares a and b and return the data that is unique to a as a difference
func diff(a []int, b []int) []int {
	tmp := make(map[int]int)
	for _, i := range b {
		tmp[i]++
	}

	diff := make([]int, 0, 0)
	for _, j := range a {
		if tmp[j] > 0 {
			tmp[j]--
			continue
		}
		diff = append(diff, j)
	}

	return diff
}

func main() {
	cases := []struct {
		expected []int
		a        []int
		b        []int
	}{
		{
			expected: []int{3},
			a:        []int{1, 2, 3},
			b:        []int{1, 2},
		},
		{
			expected: []int{},
			a:        []int{1, 2},
			b:        []int{1, 2, 3},
		},
		{
			expected: []int{2},
			a:        []int{1, 2, 2, 3},
			b:        []int{1, 2, 3},
		},
		{
			expected: []int{},
			a:        []int{1, 2, 3},
			b:        []int{1, 2, 2, 3},
		},
		{
			expected: []int{},
			a:        []int{1, 2, 3},
			b:        []int{1, 2, 3},
		},
	}

	for _, c := range cases {
		actual := diff(c.a, c.b)
		if !reflect.DeepEqual(c.expected, actual) {
			log.Fatalf("expected %#v, actual %#v", c.expected, actual)
		}
	}
}
