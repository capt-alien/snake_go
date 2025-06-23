#!/usr/bin/env python3

def longest_k_ss(s: str, k: int) -> int:
    """Return length of longest substring with ≤ k distinct chars."""
    if k == 0 or not s:
        return 0
    if k >= len(set(s)):
        return len(s)

    left = 0
    counts: dict[str, int] = {}
    longest = 0

    for right, ch in enumerate(s):
        counts[ch] = counts.get(ch, 0) + 1

        while len(counts) > k:
            counts[s[left]] -= 1
            if counts[s[left]] == 0:
                del counts[s[left]]
            left += 1

        longest = max(longest, right - left + 1)

    return longest

if __name__ == "__main__":
    assert longest_k_ss("eceba", 2) == 3
    assert longest_k_ss("aa", 1) == 2
    assert longest_k_ss("abcadcacacaca", 3) == 11
    assert longest_k_ss("", 3) == 0
    assert longest_k_ss("abc", 0) == 0
    print("All good ✔️")
