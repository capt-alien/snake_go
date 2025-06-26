package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s)
	left, right := 0, len(runes) -1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)
}

func main() {
	input := "GoLang"
	result := reverseString(input)
	fmt.Println(result) // Expected: "olleh"
}
