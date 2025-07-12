#!/usr/bin/env python3

def search_rotated_sorted_arr(nums, target):
    left = 0    #left index
    right = len(nums) -1 #right index
    found = False

    while left <= right:
        mid = (right + left)//2
        if nums[mid] == target:
            return mid

        if nums[left] <= nums[mid]: # left side sorted
            if nums[left] <= target < nums[mid]: #target betwen left and mid:
                right = mid - 1
            else:
                left = mid + 1
        else:
            if nums[mid] < target <= nums[right]:
                left = mid + 1
            else:
                right = mid -1

    return -1




if __name__ == '__main__':
    nums = [4,5,6,7,0,1,2]
    target = 0
    print(search_rotated_sorted_arr(nums, target))

