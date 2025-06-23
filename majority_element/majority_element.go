package main

import (
	"fmt"
)

func majority_element(nums []int) int {
	majority := len(nums) / 2  // integer division in Go is always floored
	tally := map[int]int{}
	for _, num := range nums {
		tally[num] = tally[num] + 1
		if tally[num] > majority {
			return num
		}
	}
	fmt.Println(tally)
	return -1
}



func main() {
	nums := []int{2, 2 ,1 ,1 ,1 ,2 ,2 }
	result := majority_element(nums)
	fmt.Println(result)
}


