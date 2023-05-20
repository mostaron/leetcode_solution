package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	times := make([]int, len(nums))
	for index, _ := range times {
		times[index] = math.MaxInt
	}
	times[0] = 0
	for i := 0; i < len(times); i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j == len(times)-1 {
				return times[i] + 1
			}
			times[i+j] = min(times[i+j], times[i]+1)
		}
	}
	return times[len(times)-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
