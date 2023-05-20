package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1}, 2))
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		} else {
			return 1
		}
	}

	for {
		i := left + right/2
		if (i == 0 && nums[i] >= target) ||
			(nums[i-1] < target && nums[i] >= target) {
			return i
		} else if i == len(nums)-1 && nums[i] < target {
			return i + 1
		}
		if nums[i] < target {
			left = i
		} else {
			right = i
		}
	}
}
