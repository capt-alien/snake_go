#!/usr/bin/env python3


def two_sum(nums, target):
    tracker = {}
    for x in range(len(nums)):
        num = nums[x]
        counter = target - num
        if counter in tracker:
            return tracker[counter], x

        tracker[num] = x

    return -1


if __name__ == '__main__':
    nums = [2, 7, 11, 15]
    target = 9
    # Expected: [0, 1]
    print(two_sum(nums, target))
