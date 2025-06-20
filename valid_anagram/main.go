package main

import (
	"fmt"
)

func main() {
	s := "anagram"
	t := "mnagraa"
	fmt.Println(isAnagram(s, t))
}

func freqMap(s string) map[rune]int {
	m := map[rune]int{}
	for _, ch := range []rune(s){
		m[ch]++
	}
	return m
}

func isAnagram(s , t string) bool {
	if len(s) != len(t){
		return false
	}
	count_s := freqMap(s)
	count_t := freqMap(t)

 	for k, v := range count_s {
 		if count_t[k] != v {
 			return false
 		}
 	}
 	return true

}
