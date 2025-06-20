package main

import (
	"fmt"
)

func main() {
	s := "abbac"
	fmt.Println(firstUnique(s))
}

func firstUnique(s string) int {
	// Your implementation goes here
	count := map[rune]int{}
	runes := []rune(s)

	for _, ch := range runes {
		count[ch]++
	}

	// run a for loop on the indexes,
	// return the first index with count 1
	for i, ch := range runes{
		if count[ch] == 1 {
			return i
		}
	}

	return -1
}
