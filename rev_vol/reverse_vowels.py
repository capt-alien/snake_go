#!/usr/bin/env python3

def reverse_vowels(s: str) -> str:
    vowels = set("aeiouAEIOU")             # faster membership test + handles case
    chars  = list(s)                       # mutable
    left, right = 0, len(chars) - 1

    while left < right:
        while left < right and chars[left] not in vowels:
            left += 1
        while left < right and chars[right] not in vowels:
            right -= 1

        chars[left], chars[right] = chars[right], chars[left]
        left  += 1
        right -= 1

    return "".join(chars)


if __name__ == '__main__':
    print(reverse_vowels("hello"))     # holle
    print(reverse_vowels("leetcode"))  # leotcede
    print(reverse_vowels("aA"))        # Aa
