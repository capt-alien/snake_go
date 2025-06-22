package main

import "fmt"

func isAlnum(b byte) bool {
    return ('A' <= b && b <= 'Z') ||
           ('a' <= b && b <= 'z') ||
           ('0' <= b && b <= '9')
}

func toLower(b byte) byte {
    if 'A' <= b && b <= 'Z' {
        return b + 32
    }
    return b
}

func isPalindrome(s string) bool {
	l , r := 0, len(s) -1
	for l < r {
		for l < r && !isAlnum(s[l]) {
			l++
		}
		for l < r && !isAlnum(s[r]) {
			r--
		}
		if l >= r {
			break
		}
		if toLower(s[l]) != toLower(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // expect true
	fmt.Println(isPalindrome("race a car"))                     // expect false
}
