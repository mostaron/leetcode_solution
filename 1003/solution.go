package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("aabcbc"))
	fmt.Println(isValid("abcabcababcc"))
	fmt.Println(isValid("abccba"))
}
func isValid(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	} else if n%3 != 0 {
		return false
	}
	i := 0
	for i <= len(s)-3 {
		if s[i:i+3] == "abc" {
			s = s[0:i] + s[i+3:]
			if i < 2 {
				i = 0
			} else {
				i -= 2
			}
		} else {
			i++
		}
	}
	return len(s) == 0
}
