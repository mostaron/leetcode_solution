package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(miceAndCheese([]int{1, 1, 3, 4}, []int{4, 4, 1, 1}, 2))
}

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	ans := 0
	diffs := make([]int, len(reward1))
	for i, o := range reward2 {
		ans += o
		diffs[i] = reward1[i] - o
	}
	sort.Ints(diffs)

	for i := 1; i <= k; i++ {
		ans += diffs[len(diffs)-i]
	}

	return ans
}
