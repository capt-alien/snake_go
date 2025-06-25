package main

import (
	"fmt"
)


func longestSubString(s string) int {
	tracker := make(map[rune]bool)
	longest := 0
	left := 0
	runes := []rune(s)

	for right := 0; right < len(runes); right++ {
		for tracker[runes[right]] {
			delete(tracker, runes[left])
			left ++
		}
		tracker[runes[right]] = true
		currentLength := right - left + 1
		if currentLength > longest {
			longest = currentLength
		}
	}
	return longest
}

func main() {
	s := "abcabcbb"
	result := longestSubString(s)
			fmt.Println(result)


}
