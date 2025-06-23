#!/usr/bin/env python3


def majority_element(nums):
    maj_line = len(nums)//2
    tally = {}
    for x in nums:
        tally[x] = tally.get(x, 0) + 1
        if tally[x] > maj_line:
            return x

    return -1


if __name__ == '__main__':
    nums = [2,2,1,1,1,2,2]
    print(majority_element(nums))
