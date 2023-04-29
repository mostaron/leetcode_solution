package main

import "fmt"

func main() {
	fmt.Println(search([]int{3, 1}, 1))

}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	i := (left + right) / 2
	for left < right {
		//fmt.Println(left, i, right)
		if nums[i] == target {
			return i
		}

		if nums[i] > target {
			if nums[left] > target && nums[i] >= nums[left] {
				left = i + 1
			} else {
				right = i - 1
			}
		} else {
			if nums[right] < target && nums[right] > nums[i] {
				right = i - 1
			} else {
				left = i + 1
			}
		}
		//
		//if nums[i] > target && nums[left] > target && nums[i] > nums[left] {
		//	fmt.Println(">")
		//	left = i + 1
		//} else {
		//	fmt.Println("<")
		//	right = i - 1
		//}
		i = (left + right) / 2

	}
	//fmt.Println(left, i, right)
	if nums[left] == target {
		return left
	}
	return -1
}
