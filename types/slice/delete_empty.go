package main

import (
	"fmt"
	"strings"
)

// delete_empty remove an empty value in slice.
func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func main() {
	s := strings.Split("/foo/bar/", "/")
	slice := delete_empty(s)

	fmt.Printf("%#v\n", s)     // []string{"", "foo", "bar", ""}
	fmt.Printf("%#v\n", slice) // []string{"foo", "bar"}
}
