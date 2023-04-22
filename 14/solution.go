package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	i := 0
	for {
		var c uint8
		for j := 0; j < len(strs); j++ {
			if len(strs[j]) == i {
				return strs[0][0:i]
			}
			if j == 0 {
				c = strs[j][i]
				continue
			}
			if strs[j][i] != c {
				return strs[0][0:i]
			}
		}
		i++
	}
}
