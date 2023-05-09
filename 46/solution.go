package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3, 4}))
}

func permute(nums []int) [][]int {
	result := [][]int{}
	result = append(result, nums[0:1])
	for i := 1; i < len(nums); i++ {
		tmp := [][]int{}
		for _, r := range result {
			tmp = append(tmp, append([]int{nums[i]}, r...))
			for j := 1; j < i; j++ {
				t := []int{}
				t = append(t, r[:j]...)
				t = append(t, nums[i])
				t = append(t, r[j:]...)
				tmp = append(tmp, t)
			}
			tmp = append(tmp, append(r, nums[i]))
		}
		result = tmp
	}
	return result
}
