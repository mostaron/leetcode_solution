package main

import "fmt"

func main() {
	fmt.Println(findMaxK([]int{-1, 2, -3, 3}))
	fmt.Println(findMaxK([]int{-1, 10, 6, 7, -7, 1}))
}

func findMaxK(nums []int) int {
	dict := make(map[int]bool)
	min, max := 1, -1
	for _, num := range nums {
		if num < 0 && num < min && dict[-num] {
			min = num
		} else if num > 0 && num > max && dict[-num] {
			max = num
		}
		dict[num] = true
	}
	if max > -min {
		return max
	}
	return -min
}
