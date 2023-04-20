package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(makeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{4, 3, 1}))
}

func printDP(dp [][]int) {
	for _, arr := range dp {
		fmt.Println(arr)
	}
}

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	sort.Ints(arr2)
	dp := make([][]int, len(arr1)+1)
	for i := 0; i < len(arr1)+1; i++ {
		dp[i] = make([]int, min(len(arr1), len(arr2))+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt
		}
	}
	dp[0][0] = -1

	for i := 1; i <= len(arr1); i++ {
		for j := 0; j <= min(i, len(arr2)); j++ {
			if arr1[i-1] > dp[i-1][j] {
				dp[i][j] = arr1[i-1]
			}
			if j > 0 && dp[i-1][j-1] != math.MaxInt {
				k := sort.SearchInts(arr2, dp[i-1][j-1]+1)
				if k < len(arr2) {
					dp[i][j] = min(dp[i][j], arr2[k])
				}
			}
			if i == len(arr1) && dp[i][j] != math.MaxInt {
				return j
			}
		}
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
