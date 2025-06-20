package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	tracker := map[int]int{}

	for x, val := range nums {
		counter := target - val

		if index, ok := tracker[counter]; ok {
		    return []int{index, x}
		}
			tracker[val] = x
		}
	return []int{-1, -1} // fallback
}

func main() {
	runTests()
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result) // Expected: [0 1]
}


func runTests() {
	tests := []struct {
		nums   []int
		target int
		expect []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{1, 2, 3}, 7, []int{-1, -1}}, // no match
	}

	for _, test := range tests {
		result := twoSum(test.nums, test.target)
		fmt.Printf("nums: %v target: %d → result: %v (expect: %v)\n",
			test.nums, test.target, result, test.expect)
	}
}
