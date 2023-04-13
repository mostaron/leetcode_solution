package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	runeArray := []rune(s)
	start := 0
	count := 0
	maxLength := 0

	for i := 0; i < len(runeArray); i++ {
		duplicatedAt := getDuplicatedAt(&runeArray, start, i)
		if duplicatedAt < 0 {
			count++
			if maxLength < count {
				maxLength = count
			}
		} else {
			start = duplicatedAt + 1
			count = i - start + 1
			//fmt.Println("!!!!!", start, maxLength)

		}
	}
	return maxLength
}
func getDuplicatedAt(l *[]rune, start int, current int) int {
	for i := start; i < current; i++ {
		if (*l)[i] == (*l)[current] {
			//fmt.Println("=====", string((*l)[i]), i, current)
			return i
		}
	}
	return -1
}
