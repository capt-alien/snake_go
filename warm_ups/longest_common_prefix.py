#!/usr/bin/env python3


def longest_common_prefix(arr):
    common = True
    cpfx = ""
    x = 0 # index counter

    # Get shortest string
    shortest = arr[0]

    for string in arr:
        if len(string) < len(shortest):
            shortest = string

    while common and x < len(shortest):
        target_letter = shortest[x]
        for s in arr:
            if s[x] != target_letter:
                common = False
                break
        if common:
            cpfx = cpfx + target_letter
        x += 1

    return cpfx


def longest_common_prefix_sorted(arr):

    arr.sort()
    first = arr[0]
    last = arr[-1]
    prefix = ""

    for i in range(min(len(first), len(last))):
        if first[i] == last[i]:
            prefix += first[i]
        else:
            break

    return prefix



if __name__ == '__main__':
    arr1 = ["flower", "flow", "flight"]
    arr2 = ["carlos", "carma", "car"]
    # print(longest_common_prefix(arr1))
    print(longest_common_prefix_sorted(arr2))
