package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

func canJump(nums []int) bool {
	jump := nums[0] - 1
	for i := 1; i < len(nums); i++ {
		if jump < 0 {
			return false
		}
		jump = max(jump, nums[i]) - 1
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
