package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("a"))
}

func longestPalindrome(s string) string {
	buff := []rune(s)
	var longest []rune
	l1, l2 := 0, 0

	for i := 0; i < len(buff); i++ {
		for j := len(buff) - 1; j >= i; j-- {
			isPair := true
			if i > l1 && j < l2 {
				break
			}
			if buff[j] != buff[i] {
				isPair = false
				continue
			}
			a, b := i, j
			for a <= b {
				if buff[a] != buff[b] {
					isPair = false
					break
				}
				a++
				b--
			}
			if isPair && len(longest) <= (j-i) {
				longest = buff[i : j+1]
				l1 = i
				l2 = j
			}
		}
	}
	return string(longest)
}
