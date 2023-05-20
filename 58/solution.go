package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("hello world  "))
}

func lengthOfLastWord(s string) int {
	count := 0
	for j := len(s) - 1; j >= 0; j-- {
		if count != 0 && s[j] == ' ' {
			return count
		}
		if s[j] != ' ' {
			count++
		}
	}
	return count
}
