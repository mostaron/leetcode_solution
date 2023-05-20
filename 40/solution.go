package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	//fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	//fmt.Println(combinationSum([]int{8, 7, 4, 3}, 11))
	//fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var rec func(candidates []int, index, target int) ([][]int, bool)
	rec = func(candidates []int, index, target int) ([][]int, bool) {
		res := [][]int{}
		cur := candidates[index]
		rest := target - cur
		if rest > 0 {

			for i := index + 1; i < len(candidates); i++ {
				if rest-candidates[index] >= 0 {
					if i > index+1 && candidates[i-1] == candidates[i] {
						continue
					}
					tmp, test := rec(candidates, i, rest)
					if test {
						for _, t := range tmp {
							res = append(res, append([]int{cur}, t...))
						}
					}
				}
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
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] <= target {
			tmp, test := rec(candidates, i, target)
			if test {
				result = append(result, tmp...)
			}
		}

	}

	return result
}
