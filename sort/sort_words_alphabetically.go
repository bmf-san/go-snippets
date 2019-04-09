package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"abc", "aba"}
	sort.Strings(s)
	fmt.Println(s) // [aba abc]
}
