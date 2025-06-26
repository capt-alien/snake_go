package main

import (
	"fmt"
	"sort"
)

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool{
		return runes[i] < runes[j]
	})
	return string(runes)
}


func groupAnagrams(strs []string) [][]string {
    tracker := map[string][]string{}

    for x :=0; x < len(strs); x++ {
    	uid := sortString(strs[x])
    	tracker[uid] = append(tracker[uid], strs[x])
    }
    var result [][]string
    for _, v := range tracker {
    	result = append(result, v)
    }
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	fmt.Println(result)
}
