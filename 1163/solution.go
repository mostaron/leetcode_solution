package main

import "fmt"

/*
**
NOT PASSED
*/
func main() {
	fmt.Println(lastSubstring("cacacb"))
}
func lastSubstring(s string) string {
	i, j := 0, 1
	for j < len(s) {
		k := 0
		for j+k < len(s) && s[i+k] == s[j+k] {
			k++
		}
		if j+k < len(s) && s[i+k] < s[j+k] {
			fmt.Println("move i from ", i, "to", j, "and j from ", j, "to ", max(j+1, i+k+1))
			i = j
			j = j + 1
		} else {
			j = j + k + 1
		}

	}
	return s[i:]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
