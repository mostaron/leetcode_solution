package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(orderlyQueue("cba", 1))
}

func orderlyQueue(s string, k int) string {
	if k == 1 {
		minStr := s
		for i := 0; i < len(s); i++ {
			s = s[1:] + s[0:1]
			minStr = min(minStr, s)
		}
		return minStr
	}
	arr := []rune(s)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return string(arr)
}

func min(s1, s2 string) string {
	for i := 0; i < len(s1); i++ {
		if s1[i] < s2[i] {
			return s1
		} else if s1[i] > s2[i] {
			return s2
		}
	}
	return s1
}
