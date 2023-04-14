package main

import "fmt"

func main() {
	queries := []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
	pattern := "FB"
	fmt.Println(camelMatch(queries, pattern))
}

func camelMatch(queries []string, pattern string) []bool {
	var result []bool

	for _, query := range queries {
		matchCount := 0
		isMatch := true
		queryArray := []rune(query)
		patternArray := []rune(pattern)
		next := 0
		for i := 0; i < len(queryArray) && isMatch; i++ {
			if next >= len(patternArray) {
				if isCapital(queryArray[i]) {
					isMatch = false
					break
				}
				continue
			}

			if patternArray[next] == queryArray[i] {
				next++
				matchCount++
				continue
			} else if isCapital(queryArray[i]) {
				isMatch = false
				break
			}
		}
		if isMatch && matchCount == len(patternArray) {
			result = append(result, true)
		} else {
			result = append(result, false)
		}

	}

	return result
}

func isCapital(pattern rune) bool {
	return pattern >= 'A' && pattern <= 'Z'
}
