package main

import "fmt"

func main() {
	fmt.Println(addNegabinary([]int{1, 1, 1, 1, 1}, []int{1, 0, 1}))
	fmt.Println(addNegabinary([]int{0}, []int{0}))
	fmt.Println(addNegabinary([]int{0}, []int{1}))
	fmt.Println(addNegabinary([]int{1, 0}, []int{1, 0}))
}

func addNegabinary(arr1 []int, arr2 []int) []int {
	m, n := len(arr1), len(arr2)
	result := []int{}
	overflow := 0
	for i := 1; i <= m || i <= n || overflow != 0; i++ {
		bit := 0
		if i <= m {
			bit += arr1[m-i]
		}
		if i <= n {
			bit += arr2[n-i]
		}
		bit -= overflow
		if bit == -1 {
			bit = 1
			overflow = -1
		} else {
			overflow = bit / 2
		}
		result = append([]int{bit % 2}, result...)
	}
	for result[0] == 0 && len(result) > 1 {
		result = result[1:]
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
