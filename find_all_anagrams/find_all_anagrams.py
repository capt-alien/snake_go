#!/usr/bin/env python3


def find_all_anagrams(s, p):
    output = []
    left = 0
    right = len(p)
    sorted_p = sorted(p)

    if len(s) < len(p):
        return output

    while right <= len(s):
        sorted_w = sorted(s[left:right])
        if sorted_w == sorted_p:
            output.append(left)

        left += 1
        right += 1

    return output


if __name__ == '__main__':
    s = "cbaebabacd"
    p = "abc"
    # Output: [0, 6]
    print(find_all_anagrams(s, p))
