package main

import (
	"fmt"
)

func main() {
	//fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	//fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	//fmt.Println(combinationSum([]int{8, 7, 4, 3}, 11))
	fmt.Println(combinationSum([]int{7, 3, 9, 6}, 6))
}

func combinationSum(candidates []int, target int) [][]int {

	var rec func(candidates []int, index, target int) ([][]int, bool)
	rec = func(candidates []int, index, target int) ([][]int, bool) {
		res := [][]int{}
		cur := candidates[index]
		rest := target - cur
		if rest > 0 {
			for index < len(candidates) {
				if rest-candidates[index] >= 0 {
					tmp, test := rec(candidates, index, rest)
					if test {
						for _, t := range tmp {
							res = append(res, append([]int{cur}, t...))
						}
					}
				}
				index++
			}
			if len(res) > 0 {
				return res, true
			} else {
				return nil, false
			}
		} else if rest == 0 {
			return [][]int{{candidates[index]}}, true
		} else {
			return nil, false
		}
	}
	result := [][]int{}
	for i := 0; i < len(candidates); i++ {
		if candidates[i] <= target {
			tmp, test := rec(candidates, i, target)
			if test {
				result = append(result, tmp...)
			}
		}

	}

	return result
}
