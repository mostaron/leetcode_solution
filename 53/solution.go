package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxVal := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maxVal = max(maxVal, dp[i])
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
