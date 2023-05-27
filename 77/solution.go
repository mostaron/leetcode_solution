package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(1, 1))
}

func combine(n int, k int) [][]int {
	pointer := make([]int, k)
	result := [][]int{}
	for i := range pointer {
		pointer[i] = i + 1
	}
	for pointer[0] <= n-k {
		//fmt.Println(pointer)
		tmp := []int{}
		tmp = append(tmp, pointer...)
		result = append(result, tmp)

		for i := 1; i <= k; i++ {
			if pointer[k-i]+1 <= n-i+1 {
				pointer[k-i]++
				for j := k - i + 1; j < k; j++ {
					pointer[j] = pointer[j-1] + 1
				}
				break
			}
		}
	}
	result = append(result, pointer)
	return result
}
