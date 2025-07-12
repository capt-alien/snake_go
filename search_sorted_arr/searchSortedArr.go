package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	left := 0
	right := len(nums) -1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		//left is sorted
		if nums[left] <= nums[mid]{
			if nums[left] <= target && target < nums[mid] {
				right = mid -1
			} else {
				left = mid + 1
			}
		} else {
			// right half sorted
			if nums[mid] < target && target <= nums[right]{
				left = mid +1
			} else {
				right = mid -1
			}
		}
	}

	return -1
}


func main(){
	testCases := []struct {
		nums	[]int
		target	int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
		{[]int{1, 2, 3}, 2},
	}

	for _, tc := range testCases {
		fmt.Printf("Searching for %d in %v → Index: %d\n", tc.target, tc.nums, search(tc.nums, tc.target))
	}
}
