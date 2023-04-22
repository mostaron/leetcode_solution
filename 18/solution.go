package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{-3, -1, 0, 2, 4, 5}, 2))
}

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for l := 0; l < len(nums)-3; l++ {
		if l != 0 && nums[l] == nums[l-1] {
			continue
		}

		for r := len(nums) - 1; l+2 < r; r-- {

			if r != len(nums)-1 && nums[r] == nums[r+1] {
				continue
			}
			rp := r - 1
			lp := l + 1
			for lp < rp {
				fmt.Println(l, lp, rp, r)
				if lp != l+1 && nums[lp] == nums[lp-1] {
					lp++
					continue
				}
				tmp := nums[l] + nums[lp] + nums[rp] + nums[r]
				if tmp < target {
					lp++
				} else if tmp > target {
					rp--
				} else {
					fmt.Println(l, lp, rp, r)
					result = append(result, []int{nums[l], nums[lp], nums[rp], nums[r]})
					lp++
					rp--
				}
			}
		}

	}
	return result
}
