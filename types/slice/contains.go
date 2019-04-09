package main

import "fmt"

// contains check whether a slice contains a certain item.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func main() {
	slice := []string{"id", "name"}

	fmt.Printf("%#v\n", contains(slice, "id"))    // true
	fmt.Printf("%#v\n", contains(slice, "name"))  // true
	fmt.Printf("%#v\n", contains(slice, "title")) // false
}
