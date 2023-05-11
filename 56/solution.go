package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 3}, {8, 10}, {2, 6}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{}
	for _, r := range intervals {
		if len(result) == 0 {
			result = append(result, r)
			continue
		}
		if r[0] <= result[len(result)-1][1] {
			result[len(result)-1][1] = max(r[1], result[len(result)-1][1])
		} else {
			result = append(result, r)
		}
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
