package main

import "sort"

func main() {

}

func kthSmallest(mat [][]int, k int) int {
	a := []int{0}
	for _, row := range mat {
		b := make([]int, 0, len(a)*len(row)) // 预分配空间
		for _, x := range a {
			for _, y := range row {
				b = append(b, x+y)
			}
		}
		sort.Ints(b)
		if len(b) > k { // 保留最小的 k 个
			b = b[:k]
		}
		a = b
	}
	return a[k-1]
}
