package main

import (
	"fmt"
)


func isValid(s string) bool {
	vMap := map[rune]rune{
		'{': '}',
        '(': ')',
        '[': ']',
	}
	tStack := [] rune {}

	for _, char := range s  {
		if _, ok := vMap[char]; ok{
			tStack = append(tStack, char)
		} else {
              if len(tStack) == 0 {
              	return false
              }
            opener := tStack[len(tStack)-1]
            tStack = tStack[:len(tStack)-1]

            if vMap[opener] != char {
            	return false
            }
		}
	}

	return true
}

func main() {
	tests := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
	for _, test := range tests {
		fmt.Printf("%s: %v\n", test, isValid(test))
	}
}
