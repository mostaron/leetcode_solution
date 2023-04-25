package main

import "fmt"

func main() {
	fmt.Println(lastSubstring(""))
}
func lastSubstring(s string) string {
	i, j := 0, 1
	for j < len(s) {
		k := 0
		for j+k < len(s) && s[i+k] == s[j+k] {
			k++
		}
		if j+k < len(s) && s[i+k] < s[j+k] {
			tmp := i
			i = j
			j = max(j+1, tmp+k+1)
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
