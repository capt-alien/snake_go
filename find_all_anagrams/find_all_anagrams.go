package main

import(
	"fmt"
	"sort"
)

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func findAllAnagrams(s string, p string) []int {
	output := []int{}
	sortP := sortString(p)
	left := 0
	right := len(p)
	if len(p) > len(s) {
		return output
	}
	for right <= len(s) {
		str := s[left:right]
		if sortString(str) == sortP {
			output = append(output, left)
		}
		right ++
		left ++
	}
	return output
}


func main() {
	s := "cbaebabacd"
	p := "abc"
	result := findAllAnagrams(s, p)
	fmt.Println(result)
}
