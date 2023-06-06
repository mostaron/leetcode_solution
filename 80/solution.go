package main

import "fmt"

func main() {
	arr := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(arr))
	fmt.Println(arr)
}

func removeDuplicates(nums []int) int {
	i, j, r := 0, 1, true

	for j < len(nums) {
		if nums[i] == nums[j] {
			if r {
				nums[i+1] = nums[j]
				i++
				j++
				r = false
			} else {
				j++
			}
		} else {
			r = true
			nums[i+1] = nums[j]
			i++
			j++
		}
	}
	return i + 1
}
