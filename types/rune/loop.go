package main

import "fmt"

func main() {
	word := "word"
	wordRune := []rune(word)
	wordLength := len(wordRune) // get a number of bytes

	for i := 0; i < wordLength; i++ {
		fmt.Print(string(wordRune[i]))
	} // word

	for i, v := range wordRune {
		fmt.Print(i, string(v))
	} // 0w1o2r3d
}
