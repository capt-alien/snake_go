package main

import (
	"fmt"
)

func lengthOfLongestKDistinct(s string, k int) int {
	if k == 0 || len(s) == 0 {
		return 0
	}
	seen := make(map[byte]struct{})
	for i := 0 ;i < len(s); i++{
		seen[s[i]] = struct{}{}
	}

	if k >= len(seen) {
		return len(s)
	}




	// ---------- 2. Core window state ----------
	// left, right  – window boundaries
	// counts       – map[byte]int (or map[rune]int) of char → frequency
	// longest      – best length seen so far
	left := 0
	counts := make(map[byte]int) // If you need full Unicode, switch to rune.

	longest := 0

	// ---------- 3. Expand the window ----------
	for right := 0; right < len(s); right++ {
		// grab current character
		ch := s[right]
		// bump its count
		counts[ch]++

		// ---------- 4. Shrink from the left while we have > k distinct ----------
		for len(counts) > k {
			leftChar := s[left]
			counts[leftChar]--
			if counts[leftChar] == 0 {
				delete(counts, leftChar)
			}
			left++ // slide window
		}

		// ---------- 5. Record best length ----------
		windowLen := right - left + 1
		if windowLen > longest {
			longest = windowLen
		}
	}

	// ---------- 6. Return result ----------
	return longest
}

func main() {
	tests := []struct {
		s   string
		k   int
		exp int
	}{
		{"eceba", 2, 3},
		{"aa", 1, 2},
		{"abcadcacacaca", 3, 11},
		{"", 3, 0},
		{"abc", 0, 0},
	}

	for _, tc := range tests {
		got := lengthOfLongestKDistinct(tc.s, tc.k)
		fmt.Printf("%q, k=%d → %d (exp %d)\n", tc.s, tc.k, got, tc.exp)
	}
}
