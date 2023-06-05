package main

import "fmt"

func main() {
	fmt.Println(applyOperations([]int{1, 2, 2, 1, 1, 0}))
	fmt.Println(applyOperations([]int{0, 1}))
}

func applyOperations(nums []int) []int {
	i := 0
	for i < len(nums)-1 {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
			i++
		}
		i++
	}
	i = 0
	j := 1
	for j < len(nums) {
		if nums[i] == 0 {
			if nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j++
			} else {
				j++
			}
		} else {
			i++
			if i == j {
				j++
			}
		}
	}
	return nums
}
