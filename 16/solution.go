package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	sum := math.MaxInt
	for l := 0; l < len(nums)-2; l++ {
		r := len(nums) - 1
		p := l + 1
		for p < r {
			tmp := nums[l] + nums[r] + nums[p]
			delta := target - tmp
			if abs(delta) < abs(target-sum) {
				fmt.Println(l, p, r, tmp)
				sum = tmp
			}
			if delta < 0 {
				r--
			} else if delta > 0 {
				p++
			} else {
				return tmp
			}
		}
	}
	return sum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
