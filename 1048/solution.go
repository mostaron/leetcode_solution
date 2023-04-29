package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isPre("xb", "xbc"))
	//fmt.Println(longestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"}))
	//fmt.Println(longestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}))
	//fmt.Println(longestStrChain([]string{"abcd", "dbqca"}))
	fmt.Println(longestStrChain([]string{"a", "ab", "ac", "bd", "abc", "abd", "abdd"}))
	fmt.Println(longestStrChain([]string{"bc", "abc", "abcd", "a", "ab"}))
}

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	fmt.Println(words)
	dp := make([]int, len(words))
	longest := 1
	for i := 0; i < len(words); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if isPre(words[j], words[i]) {
				dp[i] = max(dp[i], dp[j]+1)
				longest = max(longest, dp[i])
				fmt.Println(words[j], words[i], longest)
			}
		}
	}
	return longest
}

func isPre(pre, target string) bool {
	if len(pre)+1 != len(target) {
		return false
	}
	i, j := 0, 0
	for i < len(pre) && j < len(target) {
		if pre[i] == target[j] {
			i++
			j++
			continue
		}
		j++
	}
	return i == len(pre)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
