Reverse Vowels of a String
==========================

Problem Statement:
------------------
Write a function that takes a string and reverses only the vowels in it.
The order of the non-vowel characters should remain unchanged.

Examples:
---------
Input:  "hello"
Output: "holle"

Input:  "leetcode"
Output: "leotcede"

Input:  "aA"
Output: "Aa"

Function Signature:
-------------------
Python:
    def reverse_vowels(s: str) -> str

Go:
    func ReverseVowels(s string) string

Constraints:
------------
- Vowels are: 'a', 'e', 'i', 'o', 'u' (case-insensitive)
- Input string can contain both uppercase and lowercase letters
- All consonants and symbols must remain in their original position

Tips:
-----
- Use two pointers to swap vowels in place
- Consider using a set for quick vowel checks
- Convert string to a list (Python) or rune slice (Go) for mutation
