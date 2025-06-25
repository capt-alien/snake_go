#!/usr/bin/env python3

def longest_substring(s):
    tracker = set()
    longest = 0
    left = 0

    for right in range(len(s)):
        while s[right] in tracker:
            tracker.remove(s[left])
            left += 1

        tracker.add(s[right])
        longest = max(longest, right - left + 1)

    return longest


if __name__ == '__main__':
    s = "abcabcbb"
    print(longest_substring(s))
