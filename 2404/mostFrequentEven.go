package main

import "fmt"

func main() {
	fmt.Println(mostFrequentEven([]int{0, 0, 3, 2, 2, 1, 6, 6}))
}

func mostFrequentEven(nums []int) int {
	results := make(map[int]int)
	for _, num := range nums {
		if num%2 != 0 {
			continue
		}
		count := results[num]
		if count == 0 {
			results[num] = 1
		} else {
			results[num] += 1
		}
	}

	if len(results) == 0 {
		return -1
	}

	count := 0
	num := 0

	for k, v := range results {
		if v > count || (v == count && k < num) {
			num = k
			count = v
		}
	}
	return num
}
