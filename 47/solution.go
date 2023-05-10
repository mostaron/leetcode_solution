package main

import "fmt"

func main() {
	fmt.Println(permuteUnique([]int{3, 3, 3, 0}))
}

func permuteUnique(nums []int) [][]int {
	dict := make(map[string]bool)
	result := [][]int{}
	result = append(result, nums[0:1])
	for i := 1; i < len(nums); i++ {
		tmp := [][]int{}
		for _, r := range result {
			for j := 0; j <= i; j++ {
				t := []int{}
				t = append(t, r[:j]...)
				t = append(t, nums[i])
				t = append(t, r[j:]...)
				if !isContained(dict, t) {
					tmp = append(tmp, t)
				}

			}
		}
		result = tmp
	}
	return result
}

func isContained(dict map[string]bool, arr []int) bool {
	hash := make([]uint8, len(arr))
	for i := 0; i < len(arr); i++ {
		hash[i] = uint8(arr[i]) + '0'
	}
	if dict[string(hash)] {
		return true
	}
	dict[string(hash)] = true
	return false
}
